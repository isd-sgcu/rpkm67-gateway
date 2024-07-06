package main

import (
	"fmt"

	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	auth "github.com/isd-sgcu/rpkm67-gateway/internal/auth"
	"github.com/isd-sgcu/rpkm67-gateway/internal/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/count"
	"github.com/isd-sgcu/rpkm67-gateway/internal/db"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
	"github.com/isd-sgcu/rpkm67-gateway/internal/object"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/stamp"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"github.com/isd-sgcu/rpkm67-gateway/logger"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	objectProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/store/object/v1"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title RPKM67 API
// @version 1.0
// @description the RPKM67 API server.

// @host localhost:3001
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger := logger.New(&conf.App)
	corsHandler := config.MakeCorsConfig(conf)

	validate, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	authConn, err := grpc.NewClient(conf.Svc.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to auth service", err)
	}

	backendConn, err := grpc.NewClient(conf.Svc.Backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to backend service", err)
	}

	checkinConn, err := grpc.NewClient(conf.Svc.CheckIn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to checkin service", err)
	}

	storeConn, err := grpc.NewClient(conf.Svc.Store, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to store service", err)
	}

	authClient := authProto.NewAuthServiceClient(authConn)
	authSvc := auth.NewService(authClient, logger)
	authHdr := auth.NewHandler(authSvc, validate, logger)

	objClient := objectProto.NewObjectServiceClient(storeConn)
	objSvc := object.NewService(objClient, logger)

	userClient := userProto.NewUserServiceClient(authConn)
	userSvc := user.NewService(userClient, objSvc, logger)
	userHdr := user.NewHandler(userSvc, conf.App.MaxFileSizeMb, constant.AllowedContentType, validate, logger)

	pinClient := pinProto.NewPinServiceClient(backendConn)
	pinSvc := pin.NewService(pinClient, logger)
	pinHdr := pin.NewHandler(pinSvc, validate, logger)

	stampClient := stampProto.NewStampServiceClient(backendConn)
	stampSvc := stamp.NewService(stampClient, pinSvc, constant.PinRequiredActivity, logger)
	stampHdr := stamp.NewHandler(stampSvc, validate, logger)

	checkinClient := checkinProto.NewCheckInServiceClient(checkinConn)
	checkinSvc := checkin.NewService(checkinClient, logger)
	checkinHdr := checkin.NewHandler(checkinSvc, validate, logger)

	requestCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "Total number of API requests by path, method, status code",
	}, []string{"path", "method", "status_code"})
	requestMetrics := metrics.NewRequestMetrics(requestCounter)

	countCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "count_clicks_total",
		Help: "Total number of clicks count by name",
	}, []string{"name"})
	countMetrics := metrics.NewCountMetrics(countCounter)

	metricsReg := metrics.NewRegistry(prometheus.NewRegistry(), requestMetrics, countMetrics)
	metricsHdr := metrics.NewHandler(metricsReg, logger)

	countHdr := count.NewHandler(countMetrics, logger)

	authMiddleware := middleware.NewAuthMiddleware(authSvc, requestMetrics)

	r := router.New(conf, corsHandler, authMiddleware, requestMetrics)

	r.V1NonAuthGet("/auth/google-url", authHdr.GetGoogleLoginUrl)
	r.V1NonAuthGet("/auth/verify-google", authHdr.VerifyGoogleLogin)
	r.V1NonAuthPost("/auth/refresh", authHdr.RefreshToken)

	r.V1Get("/user/:id", userHdr.FindOne)
	r.V1Patch("/user/profile/:id", userHdr.UpdateProfile)
	r.V1Put("/user/picture/:id", userHdr.UpdatePicture)

	r.V1Post("/checkin", checkinHdr.Create)
	r.V1Get("/checkin/:userId", checkinHdr.FindByUserID)
	r.V1Get("/checkin/email/:email", checkinHdr.FindByEmail)

	r.V1Get("/stamp/:userId", stampHdr.FindByUserId)
	r.V1Post("/stamp/:userId", stampHdr.StampByUserId)

	r.V1Get("/pin", pinHdr.FindAll)
	r.V1Post("/pin/reset/:activityId", pinHdr.ResetPin)

	r.V1NonAuthPost("/count/:name", countHdr.Count)

	r.V1NonAuth.GET("/metrics", metricsHdr.ExposeMetrics)

	// Comment out in prod
	dbConn, err := db.InitDatabase(&conf.Db, conf.App.IsDevelopment())
	dbHdr := db.NewHandler(dbConn, logger)
	r.V1Get("/clean-db", dbHdr.CleanDb)

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		logger.Fatal("unable to start server")
	}
}

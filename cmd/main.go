package main

import (
	"fmt"

	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	auth "github.com/isd-sgcu/rpkm67-gateway/internal/auth"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"github.com/isd-sgcu/rpkm67-gateway/logger"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger := logger.New(&conf.App)
	corsHandler := config.MakeCorsConfig(conf)
	authMiddleware := middleware.NewAuthMiddleware(&conf.App)

	validate, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	authConn, err := grpc.NewClient(conf.Svc.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to auth service", err)
	}

	authClient := authProto.NewAuthServiceClient(authConn)
	authSvc := auth.NewService(authClient, logger)
	authHdr := auth.NewHandler(authSvc, validate, logger)

	userClient := userProto.NewUserServiceClient(authConn)
	userSvc := user.NewService(userClient, logger)
	userHdr := user.NewHandler(userSvc, conf.App.MaxFileSizeMb, constant.AllowedContentType, validate, logger)

	r := router.New(conf, corsHandler, authMiddleware)

	r.V1Get("/auth/google-url", authHdr.GetGoogleLoginUrl)
	r.V1Post("/auth/verify-google", authHdr.VerifyGoogleLogin)
	r.V1Post("/auth/test", authHdr.Test)

	r.V1Get("/user/:id", userHdr.FindOne)
	r.V1Patch("/user/profile/:id", userHdr.UpdateProfile)
	r.V1Put("/user/picture/:id", userHdr.UpdatePicture)

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		logger.Fatal("unable to start server")
	}
}

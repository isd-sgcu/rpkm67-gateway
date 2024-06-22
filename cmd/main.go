package main

import (
	"fmt"

	"github.com/isd-sgcu/rpkm67-gateway/config"
	auth "github.com/isd-sgcu/rpkm67-gateway/internal/auth"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/logger"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
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

	authConn, err := grpc.NewClient(conf.Svc.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Fatalf("cannot connect to auth service", err)
	}

	authClient := authProto.NewAuthServiceClient(authConn)
	authSvc := auth.NewService(authClient, logger)
	authHdr := auth.NewHandler(authSvc, logger)

	r := router.New(conf, corsHandler, authMiddleware)

	r.V1.GET("/auth/signin", authHdr.SignIn)
	r.V1.GET("/auth/signup", authHdr.SignUp)

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		logger.Fatal("unable to start server")
	}
}

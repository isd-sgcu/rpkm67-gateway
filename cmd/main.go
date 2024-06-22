package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/logger"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger := logger.New(&conf.App)
	corsHandler := config.MakeCorsConfig(conf)
	authMiddleware := middleware.NewAuthMiddleware(&conf.App)

	r := router.New(conf, corsHandler, authMiddleware)

	r.V1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		logger.Fatal("unable to start server")
	}
}

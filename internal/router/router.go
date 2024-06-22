package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
	V1 *gin.RouterGroup
}

func New(conf *config.Config, corsHandler config.CorsHandler, authMiddleware middleware.AuthMiddleware) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.HandlerFunc(corsHandler))
	v1 := r.Group("/api/v1")

	if conf.App.IsDevelopment() {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return &Router{r, v1}
}

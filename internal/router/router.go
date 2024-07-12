package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/config"
	_ "github.com/isd-sgcu/rpkm67-gateway/docs"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
	"github.com/isd-sgcu/rpkm67-gateway/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
)

type Router struct {
	*gin.Engine
	V1             *gin.RouterGroup
	V1NonAuth      *gin.RouterGroup
	requestMetrics metrics.RequestMetrics
	tracer         trace.Tracer
}

func New(conf *config.Config, corsHandler config.CorsHandler, authMiddleware middleware.AuthMiddleware, requestMetrics metrics.RequestMetrics, tracer trace.Tracer) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(otelgin.Middleware(conf.App.ServiceName))
	r.Use(gin.HandlerFunc(corsHandler))
	v1 := r.Group("/api/v1")
	v1.Use(gin.HandlerFunc(authMiddleware.Validate))

	v1NonAuth := r.Group("/api/v1")

	if conf.App.IsDevelopment() {
		v1NonAuth.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return &Router{r, v1, v1NonAuth, requestMetrics, tracer}
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
)

func (r *Router) V1Get(path string, handler func(c context.Ctx)) {
	r.V1.GET(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.GET, path, r.requestMetrics))
	})
}

func (r *Router) V1Post(path string, handler func(c context.Ctx)) {
	r.V1.POST(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.POST, path, r.requestMetrics))
	})
}

func (r *Router) V1Put(path string, handler func(c context.Ctx)) {
	r.V1.PUT(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.PUT, path, r.requestMetrics))
	})
}

func (r *Router) V1Patch(path string, handler func(c context.Ctx)) {
	r.V1.PATCH(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.PATCH, path, r.requestMetrics))
	})
}

func (r *Router) V1Delete(path string, handler func(c context.Ctx)) {
	r.V1.DELETE(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.DELETE, path, r.requestMetrics))
	})
}

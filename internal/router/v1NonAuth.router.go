package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
)

func (r *Router) V1NonAuthGet(path string, handler func(c context.Ctx)) {
	r.V1NonAuth.GET(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.GET, path, r.requestMetrics, r.tracer))
	})
}

func (r *Router) V1NonAuthPost(path string, handler func(c context.Ctx)) {
	r.V1NonAuth.POST(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.POST, path, r.requestMetrics, r.tracer))
	})
}

func (r *Router) V1NonAuthPut(path string, handler func(c context.Ctx)) {
	r.V1NonAuth.PUT(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.PUT, path, r.requestMetrics, r.tracer))
	})
}

func (r *Router) V1NonAuthPatch(path string, handler func(c context.Ctx)) {
	r.V1NonAuth.PATCH(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.PATCH, path, r.requestMetrics, r.tracer))
	})
}

func (r *Router) V1NonAuthDelete(path string, handler func(c context.Ctx)) {
	r.V1NonAuth.DELETE(path, func(c *gin.Context) {
		handler(context.NewContext(c, constant.DELETE, path, r.requestMetrics, r.tracer))
	})
}

package router

import "github.com/gin-gonic/gin"

func (r *Router) V1Get(path string, handler func(c Context)) {
	r.V1.GET(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Post(path string, handler func(c Context)) {
	r.V1.POST(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Put(path string, handler func(c Context)) {
	r.V1.PUT(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Patch(path string, handler func(c Context)) {
	r.V1.PATCH(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Delete(path string, handler func(c Context)) {
	r.V1.DELETE(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

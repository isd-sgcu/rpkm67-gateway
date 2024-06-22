package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/config"
)

type AuthMiddleware gin.HandlerFunc

func NewAuthMiddleware(conf *config.AppConfig) AuthMiddleware {
	return func(c *gin.Context) {
		// authHeader := c.GetHeader("Authorization")
		// if authHeader == "" {
		// 	c.Next()
		// 	return
		// }
		c.Next()
	}
}

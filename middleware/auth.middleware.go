package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/auth"
)

type AuthMiddleware interface {
	Validate(c *gin.Context)
}

type authMiddlewareImpl struct {
	authSvc auth.Service
}

func NewAuthMiddleware(authSvc auth.Service) AuthMiddleware {
	return &authMiddlewareImpl{authSvc}
}

func (m *authMiddlewareImpl) Validate(c *gin.Context) {

}

// func AuthMiddleware(conf *config.AppConfig) AuthMiddleware {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			returnError(c, apperror.UnauthorizedError("Authorization header not found"))
// 			c.Abort()
// 			return
// 		}

// 		if !strings.HasPrefix(authHeader, "Bearer ") {
// 			returnError(c, apperror.UnauthorizedError("Authorization header must start with 'Bearer'"))
// 			c.Abort()
// 			return
// 		}

// 		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
// 		token, err := auth..Validate(c, tokenString, cfg.OAuth2Config.ClientId)

// 		c.Next()
// 	}
// }

func returnError(c *gin.Context, err *apperror.AppError) {
	c.JSON(
		err.HttpCode,
		gin.H{
			"instance": c.Request.URL.Path,
			"title":    err.Id,
		},
	)
}

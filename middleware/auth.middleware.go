package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/isd-sgcu/rpkm67-gateway/internal/auth"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
)

type AuthMiddleware interface {
	Validate(c *gin.Context)
}

type authMiddlewareImpl struct {
	authSvc    auth.Service
	reqMetrics metrics.RequestMetrics
}

func NewAuthMiddleware(authSvc auth.Service, reqMetrics metrics.RequestMetrics) AuthMiddleware {
	return &authMiddlewareImpl{authSvc, reqMetrics}
}

func (m *authMiddlewareImpl) Validate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		m.returnError(c, apperror.UnauthorizedError("Authorization header not found"))
		c.Abort()
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		m.returnError(c, apperror.UnauthorizedError("Authorization header must start with 'Bearer'"))
		c.Abort()
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	res, err := m.authSvc.Validate(&dto.ValidateRequest{AccessToken: tokenString})
	if err != nil {
		m.returnError(c, apperror.UnauthorizedError("Invalid access token"))
		c.Abort()
		return
	}

	c.Set("userId", res.UserId)
	c.Set("role", res.Role)

	c.Next()
}

func (m *authMiddlewareImpl) returnError(c *gin.Context, err *apperror.AppError) {
	m.reqMetrics.AddRequest(c.Request.URL.EscapedPath(), constant.Method(c.Request.Method), err.HttpCode)
	c.JSON(
		err.HttpCode,
		gin.H{
			"error": err.Id,
		},
	)
}

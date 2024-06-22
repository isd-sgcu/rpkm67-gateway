package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Validate(c *gin.Context)
	RefreshToken(c *gin.Context)
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type handlerImpl struct {
	svc      Service
	validate validator.DtoValidator
	log      *zap.Logger
}

func NewHandler(svc Service, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:      svc,
		validate: validate,
		log:      log,
	}
}

func (h *handlerImpl) Validate(c *gin.Context) {
}

func (h *handlerImpl) RefreshToken(c *gin.Context) {
}

func (h *handlerImpl) SignUp(c *gin.Context) {
}

func (h *handlerImpl) SignIn(c *gin.Context) {
}

func (h *handlerImpl) SignOut(c *gin.Context) {
}

func (h *handlerImpl) ForgotPassword(c *gin.Context) {
}

func (h *handlerImpl) ResetPassword(c *gin.Context) {
}

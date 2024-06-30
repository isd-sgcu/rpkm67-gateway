package auth

import (
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	RefreshToken(c context.Ctx)
	GetGoogleLoginUrl(c context.Ctx)
	VerifyGoogleLogin(c context.Ctx)
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

// RefreshToken godoc
// @Summary Refresh access token with refresh token
// @Description get both new access token and refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.RefreshTokenRequest true "refresh token"
// @Success 200 {object} dto.Credential
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /auth/refresh [post]
func (h *handlerImpl) RefreshToken(c context.Ctx) {
	req := &dto.RefreshTokenRequest{}
	if err := c.Bind(req); err != nil {
		h.log.Named("auth hdr").Error("failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("auth hdr").Error("validation error", zap.Strings("errorList", errorList))
		c.BadRequestError("validation error: " + strings.Join(errorList, ","))
		return
	}

	credential, appErr := h.svc.RefreshToken(req)
	if appErr != nil {
		c.ResponseError(appErr)
		return
	}

	c.JSON(200, credential)
}

// GetGoogleLoginUrl godoc
// @Summary Get Google login url
// @Description get google login url
// @Tags auth
// @Produce json
// @Success 200 {object} dto.GetGoogleLoginUrlResponse
// @Failure 500 {object} apperror.AppError
// @Router /auth/google-url [get]
func (h *handlerImpl) GetGoogleLoginUrl(c context.Ctx) {
	res, appErr := h.svc.GetGoogleLoginUrl()
	if appErr != nil {
		c.ResponseError(appErr)
		return
	}

	c.JSON(200, res)
}

// VerifyGoogleLogin godoc
// @Summary Verify Google login
// @Description returns user's credential
// @Tags auth
// @Accept plain
// @Produce json
// @Param code path string true "Code from google login"
// @Success 200 {object} dto.VerifyGoogleLoginResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /auth/verify-google/{code} [get]
func (h *handlerImpl) VerifyGoogleLogin(c context.Ctx) {
	code := c.Query("code")
	if code == "" {
		c.BadRequestError("url parameter 'code' not found")
	}

	req := &dto.VerifyGoogleLoginRequest{
		Code: code,
	}

	credential, appErr := h.svc.VerifyGoogleLogin(req)
	if appErr != nil {
		c.ResponseError(appErr)
		return
	}

	c.JSON(200, credential)
}

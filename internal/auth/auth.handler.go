package auth

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Validate(c router.Context)
	RefreshToken(c router.Context)
	GetGoogleLoginUrl(c router.Context)
	VerifyGoogleLogin(c router.Context)
	Test(c router.Context)
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

func (h *handlerImpl) Validate(c router.Context) {
}

func (h *handlerImpl) RefreshToken(c router.Context) {
}

func (h *handlerImpl) GetGoogleLoginUrl(c router.Context) {
}

func (h *handlerImpl) VerifyGoogleLogin(c router.Context) {
}

func (h *handlerImpl) Test(c router.Context) {
	code := c.Param("code")
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

// func (h *handlerImpl) SignUp(c router.Context) {
// 	body := &dto.SignUpRequest{}
// 	if err := c.Bind(body); err != nil {
// 		h.log.Named("auth hdr").Error("failed to bind request body", zap.Error(err))
// 		c.BadRequestError(err.Error())
// 		return
// 	}

// 	if errorList := h.validate.Validate(body); errorList != nil {
// 		h.log.Named("auth hdr").Error("validation error", zap.Strings("errorList", errorList))
// 		c.BadRequestError(strings.Join(errorList, ", "))
// 		return
// 	}

// 	req := &dto.SignUpRequest{
// 		Email:     body.Email,
// 		Password:  body.Password,
// 		Firstname: body.Firstname,
// 		Lastname:  body.Lastname,
// 	}

// 	credential, appErr := h.svc.SignUp(req)
// 	if appErr != nil {
// 		c.ResponseError(appErr)
// 		return
// 	}

// 	c.JSON(201, credential)
// }

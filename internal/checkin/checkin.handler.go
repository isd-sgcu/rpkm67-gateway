package checkin

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Create(c router.Context)
	FindByEmail(c router.Context)
	FindByUserID(c router.Context)
}

func NewHandler(svc Service, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:      svc,
		validate: validate,
		log:      log,
	}
}

type handlerImpl struct {
	svc      Service
	validate validator.DtoValidator
	log      *zap.Logger
}

func (h *handlerImpl) Create(c router.Context) {
	body := &dto.CreateCheckInRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Create").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Create").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.CreateCheckInRequest{
		Email:  body.Email,
		UserID: body.UserID,
		Event:  body.Event,
	}

	res, appErr := h.svc.Create(req)
	if appErr != nil {
		h.log.Named("Create").Error("Create: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			ID:     res.CheckIn.ID,
			UserID: res.CheckIn.UserID,
			Email:  res.CheckIn.Email,
			Event:  res.CheckIn.Event,
		},
	})
}

func (h *handlerImpl) FindByEmail(c router.Context) {
	body := &dto.FindByEmailCheckInRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("FindByEmail").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("FindByEmail").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.FindByEmailCheckInRequest{
		Email: body.Email,
	}

	res, appErr := h.svc.FindByEmail(req)
	if appErr != nil {
		h.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByEmailCheckInResponse{
		CheckIns: res.CheckIns,
	})
}

func (h *handlerImpl) FindByUserID(c router.Context) {
	body := &dto.FindByUserIdCheckInRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("FindByUserID").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("FindByUserID").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.FindByUserIdCheckInRequest{
		UserID: body.UserID,
	}

	res, appErr := h.svc.FindByUserID(req)
	if appErr != nil {
		h.log.Named("FindByUserID").Error("FindByUserID: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByUserIdCheckInResponse{
		CheckIns: res.CheckIns,
	})
}

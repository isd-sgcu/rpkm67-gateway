package pin

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindAll(c context.Ctx)
	ResetPin(c context.Ctx)
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

func (h *handlerImpl) FindAll(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	req := &dto.FindAllPinRequest{}
	res, appErr := h.svc.FindAll(req)
	if appErr != nil {
		h.log.Named("FindAll").Error("FindAll: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindAllPinResponse{Pins: res.Pins})
}

func (h *handlerImpl) ResetPin(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	activityId := c.Param("workshop-id")
	if activityId == "" {
		c.BadRequestError("url parameter 'workshop-id' not found")
	}

	req := &dto.ResetPinRequest{
		ActivityId: activityId,
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("ResetPin").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.ResetPin(req)
	if appErr != nil {
		h.log.Named("ResetPin").Error("ResetPin: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.ResetPinResponse{Success: res.Success})
}

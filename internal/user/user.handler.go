package pin

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindAll(c router.Context)
	ResetPin(c router.Context)
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

func (h *handlerImpl) FindAll(c router.Context) {
	req := &dto.FindAllPinRequest{}
	res, appErr := h.svc.FindAll(req)
	if appErr != nil {
		h.log.Named("FindAll").Error("FindAll: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindAllPinResponse{Pins: res.Pins})
}

func (h *handlerImpl) ResetPin(c router.Context) {
	workshopId := c.Param("workshop-id")
	if workshopId == "" {
		c.BadRequestError("url parameter 'workshop-id' not found")
	}

	req := &dto.ResetPinRequest{
		WorkshopId: workshopId,
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

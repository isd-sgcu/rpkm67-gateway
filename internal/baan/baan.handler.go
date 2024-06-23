package baan

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindAllBaan(c router.Context)
	FindOneBaan(c router.Context)
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

func (h *handlerImpl) FindAllBaan(c router.Context) {
	req := &dto.FindAllBaanRequest{}
	res, appErr := h.svc.FindAllBaan(req)
	if appErr != nil {
		h.log.Named("FindAllBaan").Error("FindAllBaan: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindAllBaanResponse{Baans: res.Baans})
}

func (h *handlerImpl) FindOneBaan(c router.Context) {
	baanId := c.Param("id")
	if baanId == "" {
		c.BadRequestError("url parameter 'id' not found")
	}

	req := &dto.FindOneBaanRequest{
		Id: baanId,
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("FineOneBaan").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.FindOneBaan(req)
	if appErr != nil {
		h.log.Named("FindOneBaan").Error("FindOneBaan: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindOneBaanResponse{Baan: res.Baan})
}

package pin

import (
	"net/http"

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

// FindAll godoc
// @Summary Find all pins
// @Description Staff only
// @Tags pin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.FindAllPinResponse
// @Failure 403 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /pin [get]
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

// ResetPin godoc
// @Summary Reset a pin
// @Description Staff only
// @Tags pin
// @Accept plain
// @Produce json
// @Param activityId path string true "should be `workshop-1` to `workshop-5`"
// @Security BearerAuth
// @Success 200 {object} dto.ResetPinResponse
// @Failure 403 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /pin/reset/{activityId} [post]
func (h *handlerImpl) ResetPin(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	activityId := c.Param("activityId")
	if activityId == "" {
		c.BadRequestError("url parameter 'activityId' not found")
	}

	req := &dto.ResetPinRequest{
		ActivityId: activityId,
	}

	res, appErr := h.svc.ResetPin(req)
	if appErr != nil {
		h.log.Named("ResetPin").Error("ResetPin: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.ResetPinResponse{Success: res.Success})
}

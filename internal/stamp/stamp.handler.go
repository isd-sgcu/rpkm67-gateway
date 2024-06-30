package stamp

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindByUserId(c router.Context)
	StampByUserId(c router.Context)
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

// FindByUserId godoc
// @Summary Find stamp by user id
// @Description Find stamp by user id
// @Tags stamp
// @Accept plain
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dto.FindByUserIdStampResponse
// @Failure 400 {object} apperror.AppError
// @Router /stamp/{userId} [get]
func (h *handlerImpl) FindByUserId(c router.Context) {
	userId := c.Param("userId")
	if userId == "" {
		h.log.Named("FindByUserId").Error("FindByUserId: userId is empty")
		c.BadRequestError("userId should not be empty")
		return
	}

	req := &dto.FindByUserIdStampRequest{
		UserID: userId,
	}

	res, appErr := h.svc.FindByUserId(req)
	if appErr != nil {
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByUserIdStampResponse{
		Stamp: res.Stamp,
	})
}

// StampByUserId godoc
// @Summary Stamp by user id
// @Description Stamp by user id
// @Tags stamp
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body dto.StampByUserIdBodyRequest true "Stamp by user id request"
// @Success 200 {object} dto.StampByUserIdResponse
// @Failure 400 {object} apperror.AppError
// @Router /stamp/{userId} [post]
func (h *handlerImpl) StampByUserId(c router.Context) {
	userId := c.Param("userId")
	if userId == "" {
		h.log.Named("StampByUserId").Error("StampByUserId: userId is empty")
		c.BadRequestError("userId should not be empty")
		return
	}

	body := &dto.StampByUserIdBodyRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("StampByUserId").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("StampByUserId").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.StampByUserIdRequest{
		UserID:     userId,
		ActivityId: body.ActivityId,
		Pin:        body.Pin,
	}

	res, appErr := h.svc.StampByUserId(req)
	if appErr != nil {
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.StampByUserIdResponse{
		Stamp: res.Stamp,
	})
}

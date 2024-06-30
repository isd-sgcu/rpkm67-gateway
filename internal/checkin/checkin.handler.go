package checkin

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Create(c context.Ctx)
	FindByEmail(c context.Ctx)
	FindByUserID(c context.Ctx)
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

// Create godoc
// @Summary Create a check-in
// @Description Create a check-in using email, event and user_id
// @Tags checkin
// @Accept json
// @Produce json
// @Param create body dto.CreateCheckInRequest true "Create CheckIn Request"
// @Success 201 {object} dto.CreateCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin [post]
func (h *handlerImpl) Create(c context.Ctx) {
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

// FindByEmail godoc
// @Summary Find check-ins by email
// @Description Find check-ins by email
// @Tags checkin
// @Accept plain
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} dto.FindByEmailCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin/email/{email} [get]
func (h *handlerImpl) FindByEmail(c context.Ctx) {
	email := c.Param("email")
	if email == "" {
		h.log.Named("FindByEmail").Error("FindByEmail: email should not be empty")
		c.BadRequestError("email should not be empty")
		return
	}

	req := &dto.FindByEmailCheckInRequest{
		Email: email,
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

// FindByUserID godoc
// @Summary Find check-ins by user_id
// @Description Find check-ins by user_id
// @Tags checkin
// @Accept plain
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dto.FindByUserIdCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin/{userId} [get]
func (h *handlerImpl) FindByUserID(c context.Ctx) {
	userId := c.Param("userId")
	if userId == "" {
		h.log.Named("FindByUserID").Error("FindByUserID: user_id should not be empty")
		c.BadRequestError("user_id should not be empty")
		return
	}

	req := &dto.FindByUserIdCheckInRequest{
		UserID: userId,
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

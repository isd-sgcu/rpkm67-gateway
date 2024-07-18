package checkin

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
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
	userSvc  user.Service
	validate validator.DtoValidator
	log      *zap.Logger
}

func NewHandler(svc Service, userSvc user.Service, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:      svc,
		userSvc:  userSvc,
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
// @Security BearerAuth
// @Success 201 {object} dto.CreateCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin [post]
func (h *handlerImpl) Create(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	tr := c.GetTracer()
	ctx, span := tr.Start(c.RequestContext(), "handler.checkin.Create")
	defer span.End()

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

	res, appErr := h.svc.Create(ctx, req)
	if appErr != nil {
		h.log.Named("Create").Error("Create: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	resUser, appErr := h.userSvc.FindOne(&dto.FindOneUserRequest{Id: res.CheckIn.UserID})
	if appErr != nil {
		h.log.Named("Create").Error("FindOne: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}
	h.log.Info("user found", zap.Any("user", resUser.User))

	c.JSON(http.StatusCreated, &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			ID:     res.CheckIn.ID,
			UserID: res.CheckIn.UserID,
			Email:  resUser.User.Email,
			Event:  res.CheckIn.Event,
		},
		Firstname: resUser.User.Firstname,
		Lastname:  resUser.User.Lastname,
	})
}

// FindByEmail godoc
// @Summary Find check-ins by email
// @Description Find check-ins by email
// @Tags checkin
// @Accept plain
// @Produce json
// @Param email path string true "Email"
// @Security BearerAuth
// @Success 200 {object} dto.FindByEmailCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin/email/{email} [get]
func (h *handlerImpl) FindByEmail(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	tr := c.GetTracer()
	ctx, span := tr.Start(c.RequestContext(), "handler.checkin.FindByEmail")
	defer span.End()

	email := c.Param("email")
	if email == "" {
		h.log.Named("FindByEmail").Error("FindByEmail: email should not be empty")
		c.BadRequestError("email should not be empty")
		return
	}

	req := &dto.FindByEmailCheckInRequest{
		Email: email,
	}

	res, appErr := h.svc.FindByEmail(ctx, req)
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
// @Security BearerAuth
// @Success 200 {object} dto.FindByUserIdCheckInResponse
// @Failure 400 {object} apperror.AppError
// @Router /checkin/{userId} [get]
func (h *handlerImpl) FindByUserID(c context.Ctx) {
	if c.GetString("role") != "staff" {
		c.ResponseError(apperror.ForbiddenError("only staff can access this endpoint"))
		return
	}

	tr := c.GetTracer()
	ctx, span := tr.Start(c.RequestContext(), "handler.checkin.FindByUserID")
	defer span.End()

	userId := c.Param("userId")
	if userId == "" {
		h.log.Named("FindByUserID").Error("FindByUserID: user_id should not be empty")
		c.BadRequestError("user_id should not be empty")
		return
	}

	req := &dto.FindByUserIdCheckInRequest{
		UserID: userId,
	}

	res, appErr := h.svc.FindByUserID(ctx, req)
	if appErr != nil {
		h.log.Named("FindByUserID").Error("FindByUserID: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByUserIdCheckInResponse{
		CheckIns: res.CheckIns,
	})
}

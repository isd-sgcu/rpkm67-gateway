package checkin

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
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
	svc              Service
	userSvc          user.Service
	regConf          *config.RegConfig
	staffOnlyCheckin map[string]struct{}
	validate         validator.DtoValidator
	log              *zap.Logger
}

func NewHandler(svc Service, userSvc user.Service, regConf *config.RegConfig, staffOnlyCheckin map[string]struct{}, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:              svc,
		userSvc:          userSvc,
		regConf:          regConf,
		staffOnlyCheckin: staffOnlyCheckin,
		validate:         validate,
		log:              log,
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
	tr := c.GetTracer()
	ctx, span := tr.Start(c.RequestContext(), "handler.checkin.Create")
	defer span.End()

	body := &dto.CreateCheckInRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Create").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	_, isStaffOnlyCheckin := h.staffOnlyCheckin[body.Event]
	if c.GetString("role") != "staff" && isStaffOnlyCheckin {
		c.ResponseError(apperror.ForbiddenError(fmt.Sprintf("only staff can create checkin for event %s", body.Event)))
		return
	}

	ok, msg := h.checkRegTime(body.Event)
	if !ok {
		c.ForbiddenError(msg)
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Create").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	if body.StudentID != "" {
		req := &dto.FindByEmailUserRequest{Email: body.StudentID + "@student.chula.ac.th"}
		res, appErr := h.userSvc.FindByEmail(req)
		if appErr != nil {
			h.log.Named("Create").Error("FindOne: ", zap.Error(appErr))
			c.ResponseError(appErr)
			return
		}

		body.UserID = res.User.Id
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

	c.JSON(http.StatusCreated, &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			ID:          res.CheckIn.ID,
			UserID:      res.CheckIn.UserID,
			Email:       resUser.User.Email,
			Event:       res.CheckIn.Event,
			Timestamp:   res.CheckIn.Timestamp,
			IsDuplicate: res.CheckIn.IsDuplicate,
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

func (h *handlerImpl) checkRegTime(event string) (bool, string) {
	nowUTC := time.Now().UTC()
	gmtPlus7Location := time.FixedZone("GMT+7", 7*60*60)
	nowGMTPlus7 := nowUTC.In(gmtPlus7Location)
	switch event {
	case constant.RPKM_CONFIRM:
		if nowGMTPlus7.Before(h.regConf.RpkmConfirmStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: RPKM67 Confirmation Registration starts at %s", h.regConf.RpkmConfirmStart))
			return false, fmt.Sprintf("RPKM67 Confirmation Registration starts at %s", h.regConf.RpkmConfirmStart)
		}
	case constant.BAAN_RESULT:
		if nowGMTPlus7.Before(h.regConf.BaanResultStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: Baan Selection Result starts at %s", h.regConf.BaanResultStart))
			return false, fmt.Sprintf("Baan Selection Result starts at %s", h.regConf.BaanResultStart)
		}
	case constant.RPKM_DAY_ONE:
		if nowGMTPlus7.Before(h.regConf.RpkmDayOneStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: RPKM67 Day One Registration starts at %s", h.regConf.RpkmDayOneStart))
			return false, fmt.Sprintf("RPKM67 Day One Registration starts at %s", h.regConf.RpkmDayOneStart)
		}
	case constant.RPKM_DAY_TWO:
		if nowGMTPlus7.Before(h.regConf.RpkmDayTwoStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: RPKM67 Day Two Registration starts at %s", h.regConf.RpkmDayTwoStart))
			return false, fmt.Sprintf("RPKM67 Day Two Registration starts at %s", h.regConf.RpkmDayTwoStart)
		}
	case constant.FRESHY_NIGHT_CONFIRM:
		if nowGMTPlus7.Before(h.regConf.FreshyNightConfirmStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: Freshy Night Confirmation Registration starts at %s", h.regConf.FreshyNightConfirmStart))
			return false, fmt.Sprintf("Freshy Night Confirmation Registration starts at %s", h.regConf.FreshyNightConfirmStart)
		} else if nowGMTPlus7.After(h.regConf.FreshyNightConfirmEnd) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: Freshy Night Confirmation Registration ends at %s", h.regConf.FreshyNightConfirmEnd))
			return false, fmt.Sprintf("Freshy Night Confirmation Registration ends at %s", h.regConf.FreshyNightConfirmEnd)
		}
	case constant.FRESHY_NIGHT:
		if nowGMTPlus7.Before(h.regConf.FreshyNightStart) {
			h.log.Named("checkRegTime").Warn(fmt.Sprintf("Forbidden: Freshy Night Registration starts at %s", h.regConf.FreshyNightStart))
			return false, fmt.Sprintf("Freshy Night Registration starts at %s", h.regConf.FreshyNightStart)
		}
	default:
		h.log.Named("checkRegTime").Warn("Forbidden: Invalid event")
		return false, "Invalid event name"
	}

	return true, ""
}

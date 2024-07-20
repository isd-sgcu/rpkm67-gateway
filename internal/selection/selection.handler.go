package selection

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/group"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Create(c context.Ctx)
	FindByGroupId(c context.Ctx)
	Update(c context.Ctx)
	Delete(c context.Ctx)
	CountByBaanId(c context.Ctx)
}

type handlerImpl struct {
	svc      Service
	groupSvc group.Service
	validate validator.DtoValidator
	rpkmConf *config.RpkmConfig
	log      *zap.Logger
}

func NewHandler(svc Service, groupSvc group.Service, rpkmConf *config.RpkmConfig, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:      svc,
		groupSvc: groupSvc,
		validate: validate,
		rpkmConf: rpkmConf,
		log:      log,
	}
}

// Create godoc
// @Summary Create selection, only group leader can call
// @Description used when creating a selection on UNSELECTED order and baan
// @Tags selection
// @Accept json
// @Produce json
// @Param body body dto.CreateSelectionRequest true "create selection request"
// @Security BearerAuth
// @Success 200 {object} dto.UpdateConfirmGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 403 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /selection [post]
func (h *handlerImpl) Create(c context.Ctx) {
	h.checkGroupLeader(c)

	body := &dto.CreateSelectionRequest{}
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

	res, appErr := h.svc.Create(body)
	if appErr != nil {
		h.log.Named("Create").Error("Create: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateSelectionResponse{Selection: res.Selection})
}

// FindByGroupId godoc
// @Summary find selection by group id
// @Description used when getting all selections in a group
// @Tags selection
// @Accept json
// @Produce json
// @Param groupId path string true "groupId of request sender"
// @Security BearerAuth
// @Success 200 {object} dto.FindByGroupIdSelectionResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /selection/{groupId} [get]
func (h *handlerImpl) FindByGroupId(c context.Ctx) {
	groupId := c.Param("groupId")
	if groupId == "" {
		h.log.Named("FindByGroupIdSelection").Error("Param: groupId not found")
		c.BadRequestError("url parameter 'groupId' not found")
		return
	}

	req := &dto.FindByGroupIdSelectionRequest{
		GroupId: groupId,
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("FindByGroupIdSelection").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.FindByGroupId(req)
	if appErr != nil {
		h.log.Named("FindByGroupIdSelection").Error("FindByGroupIdSelection: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByGroupIdSelectionResponse{Selections: res.Selections})
}

// Update godoc
// @Summary Update selection, only group leader can call
// @Description used when selecting a selection on SELECTED order or baan
// @Tags selection
// @Accept json
// @Produce json
// @Param body body dto.UpdateSelectionRequest true "update selection request"
// @Security BearerAuth
// @Success 200 {object} dto.UpdateSelectionResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 403 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /selection [patch]
func (h *handlerImpl) Update(c context.Ctx) {
	h.checkGroupLeader(c)

	body := &dto.UpdateSelectionRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Update").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Update").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.Update(body)
	if appErr != nil {
		h.log.Named("Update").Error("Update: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateSelectionResponse{
		Selection: res.Selection,
	})
}

// Delete godoc
// @Summary Delete selection, only group leader can call
// @Description used when removing a selection from the list
// @Tags selection
// @Accept json
// @Produce json
// @Param body body dto.DeleteSelectionRequest true "delete selection request"
// @Security BearerAuth
// @Success 200 {object} dto.DeleteSelectionResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 403 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /selection [delete]
func (h *handlerImpl) Delete(c context.Ctx) {
	h.checkGroupLeader(c)

	body := &dto.DeleteSelectionRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Delete").Error("Bind: ", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Delete").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.Delete(body)
	if appErr != nil {
		h.log.Named("Delete").Error("Delete: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.DeleteSelectionResponse{
		Success: res.Success,
	})
}

// CountByBaanId godoc
// @Summary count selections by baan id
// @Description displayed in baan selection page to show how many people are interested in each baan
// @Tags selection
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.CountByBaanIdSelectionResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /selection/count-by-baan [get]
func (h *handlerImpl) CountByBaanId(c context.Ctx) {
	res, appErr := h.svc.CountByBaanId()
	if appErr != nil {
		h.log.Named("CountByBaanId").Error("CountByBaanId: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.CountByBaanIdSelectionResponse{
		BaanCounts: res.BaanCounts,
	})
}

func (h *handlerImpl) checkGroupLeader(c context.Ctx) {
	userId := c.GetString("userId")

	res, err := h.groupSvc.FindByUserId(&dto.FindByUserIdGroupRequest{UserId: userId})
	if err != nil {
		h.log.Named("checkGroupLeader").Error("FindByUserId: ", zap.Error(err))
		c.InternalServerError("Cannot get user's group")
		c.Abort()
		return
	}

	if res.Group.LeaderID != userId {
		h.log.Named("checkGroupLeader").Error("Forbidden: You are not the leader of this group")
		c.ForbiddenError("You are not the leader of this group")
		c.Abort()
		return
	}

	c.Next()
}

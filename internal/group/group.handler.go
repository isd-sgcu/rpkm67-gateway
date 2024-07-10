package group

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindByUserId(c context.Ctx)
	FindByToken(c context.Ctx)
	UpdateConfirm(c context.Ctx)
	Join(c context.Ctx)
	Leave(c context.Ctx)
	SwitchGroup(c context.Ctx) // basically leave current group and join another group
	DeleteMember(c context.Ctx)
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

// FindByUserId godoc
// @Summary Find group by user id
// @Description user must be member of that group
// @Tags group
// @Accept plain
// @Produce json
// @Param userId path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} dto.FindByUserIdGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/{userId} [get]
func (h *handlerImpl) FindByUserId(c context.Ctx) {
	userId := c.Param("userId")
	if userId == "" {
		c.BadRequestError("url parameter 'user_id' not found")
	}

	req := &dto.FindByUserIdGroupRequest{
		UserId: userId,
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("FindByUserId").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.FindByUserId(req)
	if appErr != nil {
		h.log.Named("FindByUserId").Error("FindByUserId: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByUserIdGroupResponse{
		Group: res.Group,
	})
}

// FindByToken godoc
// @Summary Find group by group's token
// @Description Group leader invites member by giving them the token
// @Tags group
// @Accept plain
// @Produce json
// @Param token query string true "token of group"
// @Security BearerAuth
// @Success 200 {object} dto.FindByTokenGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/token [get]
func (h *handlerImpl) FindByToken(c context.Ctx) {
	token := c.Query("token")
	if token == "" {
		c.BadRequestError("url parameter 'token' not found")
	}

	req := &dto.FindByTokenGroupRequest{
		Token: token,
	}

	res, appErr := h.svc.FindByToken(req)
	if appErr != nil {
		h.log.Named("FindByToken").Error("FindByToken: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByTokenGroupResponse{
		Id:     res.Id,
		Token:  res.Token,
		Leader: res.Leader,
	})
}

// UpdateConfirm godoc
// @Summary Update group isConfirmed status
// @Description only group leader can update this status
// @Tags group
// @Accept json
// @Produce json
// @Param userId path string true "userId of request sender (must be group leader)"
// @Param body body dto.UpdateConfirmGroupBody true "update confirm request"
// @Security BearerAuth
// @Success 200 {object} dto.UpdateConfirmGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/{userId} [put]
func (h *handlerImpl) UpdateConfirm(c context.Ctx) {
	userId := c.Param("userId")
	if userId == "" {
		c.BadRequestError("url parameter 'user_id' not found")
	}

	body := &dto.UpdateConfirmGroupBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("UpdateConfirm").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("UpdateConfirm").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.UpdateConfirmGroupRequest{
		LeaderId:    userId,
		IsConfirmed: body.IsConfirmed,
	}

	res, appErr := h.svc.UpdateConfirm(req)
	if appErr != nil {
		h.log.Named("UpdateConfirm").Error("Update: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateConfirmGroupResponse{
		Group: res.Group,
	})
}

// Join godoc
// @Summary User joins another group
// @Description previous group gets deleted, cannot join another group you are the leader of group with more than 1 member
// @Tags group
// @Accept json
// @Produce json
// @Param body body dto.JoinGroupRequest true "join request"
// @Security BearerAuth
// @Success 200 {object} dto.JoinGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/join [post]
func (h *handlerImpl) Join(c context.Ctx) {
	body := &dto.JoinGroupRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Join").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Join").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.JoinGroupRequest{
		Token:  body.Token,
		UserId: body.UserId,
	}

	res, appErr := h.svc.Join(req)
	if appErr != nil {
		h.log.Named("Join").Error("Join: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.JoinGroupResponse{
		Group: res.Group,
	})
}

// Leave godoc
// @Summary User leaves group
// @Description cannot leave group if you are leader of a group (but you can use `join` to join another group)
// @Tags group
// @Accept json
// @Produce json
// @Param body body dto.LeaveGroupRequest true "leave request"
// @Security BearerAuth
// @Success 200 {object} dto.LeaveGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/leave [post]
func (h *handlerImpl) Leave(c context.Ctx) {
	body := &dto.LeaveGroupRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Leave").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Leave").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.LeaveGroupRequest{
		UserId: body.UserId,
	}

	res, appErr := h.svc.Leave(req)
	if appErr != nil {
		h.log.Named("Leave").Error("Leave: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.LeaveGroupResponse{
		Group: res.Group,
	})
}

func (h *handlerImpl) SwitchGroup(c context.Ctx) {
	body := &dto.SwitchGroupBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("SwitchGroup").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("SwitchGroup").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	leaveReq := &dto.LeaveGroupRequest{
		UserId: body.UserId,
	}
	_, appErr := h.svc.Leave(leaveReq)
	if appErr != nil {
		h.log.Named("SwitchGroup").Error("Leave: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	joinReq := &dto.JoinGroupRequest{
		Token:  body.NewGroupToken,
		UserId: body.UserId,
	}
	res, appErr := h.svc.Join(joinReq)
	if appErr != nil {
		h.log.Named("SwitchGroup").Error("Join: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.SwitchGroupResponse{
		Group: res.Group,
	})
}

// DeleteMember godoc
// @Summary Removes user from group
// @Description only leader can remove member, cannot remove yourself
// @Tags group
// @Accept json
// @Produce json
// @Param body body dto.DeleteMemberGroupBody true "delete request"
// @Security BearerAuth
// @Success 200 {object} dto.DeleteMemberGroupResponse
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /group/delete-member [delete]
func (h *handlerImpl) DeleteMember(c context.Ctx) {
	body := &dto.DeleteMemberGroupBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("DeleteMember").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("DeleteMember").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.DeleteMemberGroupRequest{
		UserId:   body.DeletedUserId,
		LeaderId: body.RequestingUserId,
	}

	res, appErr := h.svc.DeleteMember(req)
	if appErr != nil {
		h.log.Named("DeleteMember").Error("DeleteMember: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.DeleteMemberGroupResponse{
		Group: res.Group,
	})
}

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

func (h *handlerImpl) FindByUserId(c context.Ctx) {
	userId := c.Param("id")
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

func (h *handlerImpl) FindByToken(c context.Ctx) {
	body := &dto.FindByTokenGroupRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("FindByToken").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("FindByToken").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.FindByTokenGroupRequest{
		Token: body.Token,
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

func (h *handlerImpl) UpdateConfirm(c context.Ctx) {
	body := &dto.UpdateConfirmGroupRequest{}
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
		// Group: body.Group,
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
	body := &dto.LeaveGroupRequest{}
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

	req := &dto.LeaveGroupRequest{
		UserId: body.UserId,
	}

	res, appErr := h.svc.Leave(req)
	if appErr != nil {
		h.log.Named("SwitchGroup").Error("Leave: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.LeaveGroupResponse{
		Group: res.Group,
	})
}

func (h *handlerImpl) DeleteMember(c context.Ctx) {
	body := &dto.DeleteMemberGroupRequest{}
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
		UserId:   body.UserId,
		LeaderId: body.LeaderId,
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

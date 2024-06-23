package group

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	FindOne(c router.Context)
	FindByToken(c router.Context)
	Update(c router.Context)
	Join(c router.Context)
	DeleteMember(c router.Context)
	Leave(c router.Context)
	SelectBaan(c router.Context)
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

func (h *handlerImpl) DeleteMember(c router.Context) {
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

func (h *handlerImpl) FindByToken(c router.Context) {
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

func (h *handlerImpl) FindOne(c router.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.BadRequestError("url parameter 'user_id' not found")
	}

	req := &dto.FindOneGroupRequest{
		UserId: userId,
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		h.log.Named("FindOne").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.FindOne(req)
	if appErr != nil {
		h.log.Named("FindOne").Error("FindOne: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindOneGroupResponse{
		Group: res.Group,
	})
}

func (h *handlerImpl) Join(c router.Context) {
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

func (h *handlerImpl) Leave(c router.Context) {
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

func (h *handlerImpl) SelectBaan(c router.Context) {
	body := &dto.SelectBaanRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("SelectBaan").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("SelectBaan").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := &dto.SelectBaanRequest{
		UserId: body.UserId,
		Baans:  body.Baans,
	}

	res, appErr := h.svc.SelectBaan(req)
	if appErr != nil {
		h.log.Named("SelectBaan").Error("SelectBaan: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.SelectBaanResponse{
		Success: res.Success,
	})
}

func (h *handlerImpl) Update(c router.Context) {
	body := &dto.UpdateGroupRequest{}
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

	req := &dto.UpdateGroupRequest{
		Group: body.Group,
	}

	res, appErr := h.svc.Update(req)
	if appErr != nil {
		h.log.Named("Update").Error("Update: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateGroupResponse{
		Group: res.Group,
	})
}

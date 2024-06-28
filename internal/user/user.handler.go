package user

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
	Update(c router.Context)
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

func (h *handlerImpl) FindOne(c router.Context) {
	id := c.Param("id")
	if id == "" {
		h.log.Named("FindOne").Error("Param: id not found")
		c.BadRequestError("url parameter 'id' not found")
		return
	}

	req := &dto.FindOneUserRequest{
		Id: id,
	}

	res, appErr := h.svc.FindOne(req)
	if appErr != nil {
		h.log.Named("FindOne").Error("FindOne: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindOneUserResponse{User: res.User})
}

func (h *handlerImpl) Update(c router.Context) {
	id := c.Param("id")
	if id == "" {
		h.log.Named("Update").Error("Param: id not found")
		c.BadRequestError("url parameter 'id' not found")
		return
	}

	body := &dto.UpdateUserRequestBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("Update").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError("Bind: failed to bind request body")
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("Update").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := h.createUpdateUserRequestDto(id, body)

	res, appErr := h.svc.Update(req)
	if appErr != nil {
		h.log.Named("Update").Error("Update: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateUserResponse{User: res.User})
}

func (h *handlerImpl) createUpdateUserRequestDto(id string, body *dto.UpdateUserRequestBody) *dto.UpdateUserRequest {
	return &dto.UpdateUserRequest{
		Id:          id,
		Nickname:    body.Nickname,
		Title:       body.Title,
		Firstname:   body.Firstname,
		Lastname:    body.Lastname,
		Year:        body.Year,
		Faculty:     body.Faculty,
		Tel:         body.Tel,
		ParentTel:   body.ParentTel,
		Parent:      body.Parent,
		FoodAllergy: body.FoodAllergy,
		DrugAllergy: body.DrugAllergy,
		Illness:     body.Illness,
		PhotoKey:    body.PhotoKey,
		PhotoUrl:    body.PhotoUrl,
		Baan:        body.Baan,
		ReceiveGift: body.ReceiveGift,
		GroupId:     body.GroupId,
	}
}

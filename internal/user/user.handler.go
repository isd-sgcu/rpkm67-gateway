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
	UpdateProfile(c router.Context)
	UpdatePicture(c router.Context)
}

type handlerImpl struct {
	svc                Service
	maxFileSize        int
	allowedContentType map[string]struct{}
	validate           validator.DtoValidator
	log                *zap.Logger
}

func NewHandler(svc Service, maxFileSize int, allowedContentType map[string]struct{}, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:                svc,
		maxFileSize:        maxFileSize,
		allowedContentType: allowedContentType,
		validate:           validate,
		log:                log,
	}
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

func (h *handlerImpl) UpdateProfile(c router.Context) {
	id := c.Param("id")
	if id == "" {
		h.log.Named("UpdateProfile").Error("Param: id not found")
		c.BadRequestError("url parameter 'id' not found")
		return
	}

	body := &dto.UpdateUserProfileBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("UpdateProfile").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError("Bind: failed to bind request body")
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("UpdateProfile").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	req := h.createUpdateUserRequestDto(id, body)

	res, appErr := h.svc.UpdateProfile(req)
	if appErr != nil {
		h.log.Named("UpdateProfile").Error("Update: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateUserProfileResponse{User: res.User})
}

func (h *handlerImpl) UpdatePicture(c router.Context) {
	id := c.Param("id")
	if id == "" {
		h.log.Named("UpdatePicture").Error("Param: id not found")
		c.BadRequestError("url parameter 'id' not found")
		return
	}

	file, err := c.FormFile("picture", h.allowedContentType, int64(h.maxFileSize))
	if err != nil {
		h.log.Named("UpdatePicture").Error("FormFile: failed to get file", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	req := &dto.UpdateUserPictureRequest{
		Id:   id,
		File: file,
	}

	res, appErr := h.svc.UpdatePicture(req)
	if appErr != nil {
		h.log.Named("UpdatePicture").Error("UpdatePicture: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.UpdateUserPictureResponse{User: res.User})
}

func (h *handlerImpl) createUpdateUserRequestDto(id string, body *dto.UpdateUserProfileBody) *dto.UpdateUserProfileRequest {
	return &dto.UpdateUserProfileRequest{
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
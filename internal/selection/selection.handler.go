package selection

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"github.com/isd-sgcu/rpkm67-gateway/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	CreateSelection(c router.Context)
	FindByGroupIdSelection(c router.Context)
	DeleteSelection(c router.Context)
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

func (h *handlerImpl) CreateSelection(c router.Context) {
	body := &dto.CreateSelectionRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("CreateSelection").Error("Bind: failed to bind request body", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("CreateSelection").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.CreateSelection(body)
	if appErr != nil {
		h.log.Named("CreateSelection").Error("CreateSelection: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateSelectionResponse{Selection: res.Selection})
}

func (h *handlerImpl) FindByGroupIdSelection(c router.Context) {
	groupId := c.Param("id")
	if groupId == "" {
		h.log.Named("FindByGroupIdSelection").Error("Param: id not found")
		c.BadRequestError("url parameter 'id' not found")
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

	res, appErr := h.svc.FindByGroupIdSelection(req)
	if appErr != nil {
		h.log.Named("FindByGroupIdSelection").Error("FindByGroupIdSelection: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.FindByGroupIdSelectionResponse{
		Selections: res.Selections,
	})
}

func (h *handlerImpl) DeleteSelection(c router.Context) {
	body := &dto.DeleteSelectionRequest{}
	if err := c.Bind(body); err != nil {
		h.log.Named("UpdateSelection").Error("Bind: ", zap.Error(err))
		c.BadRequestError(err.Error())
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("UpdateSelection").Error("Validate: ", zap.Strings("errorList", errorList))
		c.BadRequestError(strings.Join(errorList, ", "))
		return
	}

	res, appErr := h.svc.DeleteSelection(body)
	if appErr != nil {
		h.log.Named("UpdateSelection").Error("UpdateSelection: ", zap.Error(appErr))
		c.ResponseError(appErr)
		return
	}

	c.JSON(http.StatusOK, &dto.DeleteSelectionResponse{
		Success: res.Success,
	})
}

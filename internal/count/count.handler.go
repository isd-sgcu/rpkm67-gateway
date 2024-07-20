package count

import (
	"net/http"

	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
	"go.uber.org/zap"
)

type Handler interface {
	Count(c context.Ctx)
}

type handlerImpl struct {
	svc          Service
	countMetrics metrics.CountMetrics
	log          *zap.Logger
}

func NewHandler(svc Service, countMetrics metrics.CountMetrics, log *zap.Logger) Handler {
	return &handlerImpl{
		countMetrics: countMetrics,
		log:          log,
	}
}

// Upload godoc
// @Summary Count clicks
// @Description Add 1 to count metrics by name
// @Tags count
// @Accept json
// @Produce json
// @Param name path string true "Name of the count metric"
// @Success 201 {object} dto.CountResponse
// @Failure 400 {object} apperror.AppError
// @Router /count/{name} [post]
func (h *handlerImpl) Count(c context.Ctx) {
	name := c.Param("name")
	if name == "" {
		h.log.Named("Count").Error("Param: url parameter 'name' not found")
		c.BadRequestError("url parameter 'name' not found")
		return
	}

	h.countMetrics.Increment(name)

	res, err := h.svc.Create(&dto.CreateCountRequest{Name: name})
	if err != nil {
		h.log.Named("Count").Error("Create: failed to create count", zap.Error(err))
		c.InternalServerError(err.Error())
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateCountResponse{
		Count: res.Count,
	})
}

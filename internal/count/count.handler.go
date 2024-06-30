package count

import (
	"net/http"

	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
	"github.com/isd-sgcu/rpkm67-gateway/internal/router"
	"go.uber.org/zap"
)

type Handler interface {
	Count(c router.Context)
}

type handlerImpl struct {
	countMetrics metrics.CountMetrics
	log          *zap.Logger
}

func NewHandler(countMetrics metrics.CountMetrics, log *zap.Logger) Handler {
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
// @Router /count [post]
func (h *handlerImpl) Count(c router.Context) {
	name := c.Param("name")
	if name == "" {
		h.log.Named("Count").Error("Param: url parameter 'name' not found")
		c.BadRequestError("url parameter 'name' not found")
		return
	}

	h.countMetrics.Increment(name)

	c.JSON(http.StatusCreated, &dto.CountResponse{
		Success: true,
	})
}

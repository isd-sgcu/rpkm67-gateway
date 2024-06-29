package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type Handler interface {
	ExposeMetrics(c *gin.Context)
}

type handlerImpl struct {
	reg Registry
	log *zap.Logger
}

func NewHandler(reg Registry, log *zap.Logger) Handler {
	return &handlerImpl{
		reg: reg,
		log: log,
	}
}

func (h *handlerImpl) ExposeMetrics(c *gin.Context) {
	promhttp.HandlerFor(h.reg.Registry(), promhttp.HandlerOpts{EnableOpenMetrics: true}).
		ServeHTTP(c.Writer, c.Request)
}

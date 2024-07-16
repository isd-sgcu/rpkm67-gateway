package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Registry interface {
	Registry() *prometheus.Registry
}

type registryImpl struct {
	reg *prometheus.Registry
}

func NewRegistry(registry *prometheus.Registry, requestMetrics RequestMetrics, countMetrics CountMetrics) Registry {
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		requestMetrics.GetCounterVec(),
		countMetrics.GetCounterVec(),
	)

	return &registryImpl{
		reg: registry,
	}
}

func (m *registryImpl) Registry() *prometheus.Registry {
	return m.reg
}

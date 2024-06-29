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

func NewRegistry(registry *prometheus.Registry) Registry {
	// requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
	// 	Name:    "http_request_duration_seconds",
	// 	Help:    "A histogram of the HTTP request durations in seconds.",
	// 	Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	// })

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		// requestsMetrics.GetCounterVec(),
		// requestDurations,
	)

	return &registryImpl{
		reg: registry,
	}
}

func (m *registryImpl) Registry() *prometheus.Registry {
	return m.reg
}

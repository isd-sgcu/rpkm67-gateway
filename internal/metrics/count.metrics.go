package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type CountMetrics interface {
	Increment(name string)
	GetCounterVec() *prometheus.CounterVec
}

type countMetricsImpl struct {
	countCounter *prometheus.CounterVec
}

func NewCountMetrics(countCounter *prometheus.CounterVec) CountMetrics {
	return &countMetricsImpl{
		countCounter: countCounter,
	}
}

func (m *countMetricsImpl) Increment(name string) {
	m.countCounter.WithLabelValues(name).Inc()
}

func (m *countMetricsImpl) GetCounterVec() *prometheus.CounterVec {
	return m.countCounter
}

package metrics

import (
	"fmt"

	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/prometheus/client_golang/prometheus"
)

type RequestMetrics interface {
	AddRequest(domain constant.Domain, method constant.Method, statusCode int, duration int)
	GetCounterVec() *prometheus.CounterVec
}

type requestMetricsImpl struct {
	requestCounter *prometheus.CounterVec
}

func NewRequestMetrics(requestCounter *prometheus.CounterVec) RequestMetrics {
	return &requestMetricsImpl{
		requestCounter: requestCounter,
	}
}

func (m *requestMetricsImpl) AddRequest(domain constant.Domain, method constant.Method, statusCode int, duration int) {
	m.requestCounter.WithLabelValues(
		domain.String(), method.String(), fmt.Sprint(statusCode), fmt.Sprint(duration)).Inc()
}

func (m *requestMetricsImpl) GetCounterVec() *prometheus.CounterVec {
	return m.requestCounter
}

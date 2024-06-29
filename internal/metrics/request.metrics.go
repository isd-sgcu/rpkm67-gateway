package metrics

import (
	"fmt"

	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/prometheus/client_golang/prometheus"
)

type RequestMetrics interface {
	AddRequest(path string, method constant.Method, statusCode int)
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

func (m *requestMetricsImpl) AddRequest(path string, method constant.Method, statusCode int) {
	m.requestCounter.WithLabelValues(
		path, method.String(), fmt.Sprint(statusCode)).Inc()
}

func (m *requestMetricsImpl) GetCounterVec() *prometheus.CounterVec {
	return m.requestCounter
}

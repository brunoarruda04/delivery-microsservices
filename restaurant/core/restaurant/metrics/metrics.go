package metrics

import (
	"restaurant/core/restaurant/domain"
	"restaurant/core/restaurant/service"
	"time"

	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	requestCount   *prometheus.Counter
	requestLatency *prometheus.Summary
	next           service.Service
}

func NewMetricsMiddleware(next service.Service) service.Service {
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "restaurant_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, []string{"method"})

	requestLatency := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "restaurant_service",
		Name:      "request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, []string{"method"})

	return &metricsMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           next,
	}
}

func (mw *metricsMiddleware) Create(restaurant domain.Restaurant) (string, error) {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Create").Add(1)
		mw.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Create(restaurant)
}

func (mw *metricsMiddleware) Get(id string) (domain.Restaurant, error) {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Get").Add(1)
		mw.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Get(id)
}

func (mw *metricsMiddleware) Update(id string, restaurant domain.Restaurant) error {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Update").Add(1)
		mw.requestLatency.With("method", "Update").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Update(id, restaurant)
}

func (mw *metricsMiddleware) Delete(id string) error {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Delete").Add(1)
		mw.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Delete(id)
}

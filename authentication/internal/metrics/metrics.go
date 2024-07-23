package metrics

import (
	"authentication/internal/domain"
	"authentication/internal/service"
	"time"

	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	requestCount   *prometheus.Counter
	requestLatency *prometheus.Summary
	next           service.AuthService
}

func NewMetricsMiddleware(next service.AuthService) service.AuthService {
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "authentication_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, []string{"method"})

	requestLatency := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "authentication_service",
		Name:      "request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, []string{"method"})

	return &metricsMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           next,
	}
}

func (mw *metricsMiddleware) Login(username, password string) (string, error) {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Login").Add(1)
		mw.requestLatency.With("method", "Login").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Login(username, password)
}

func (mw *metricsMiddleware) Register(username, password string, role domain.Role) (string, error) {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Register").Add(1)
		mw.requestLatency.With("method", "Register").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.Register(username, password, role)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"restaurant/internal/restaurant/endpoint"
	"restaurant/internal/restaurant/metrics"
	"restaurant/internal/restaurant/middleware"
	"restaurant/internal/restaurant/repository"
	"restaurant/internal/restaurant/service"
	"restaurant/internal/restaurant/tracing"
	"restaurant/internal/restaurant/transport"
	"syscall"

	"github.com/go-kit/log"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	tracer, closer := tracing.InitTracer("restaurant-service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	repo, err := repository.NewPostgresRepository()
	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}

	svc := service.NewService(repo)
	svc = middleware.NewLoggingMiddleware(logger, svc)
	svc = tracing.NewTracingMiddleware(tracer, svc)
	svc = metrics.NewMetricsMiddleware(svc)

	endpoints := endpoint.MakeEndpoints(svc)
	handler := transport.MakeHTTPHandler(endpoints)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "HTTP", "addr", ":8080")
		errs <- http.ListenAndServe(":8080", handler)
	}()
	go func() {
		logger.Log("transport", "HTTP", "addr", ":8081")
		http.Handle("/metrics", promhttp.Handler())
		errs <- http.ListenAndServe(":8081", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errs)
}

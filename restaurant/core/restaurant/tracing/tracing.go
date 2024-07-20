package tracing

import (
	"io"
	"log"
	"restaurant/core/restaurant/domain"
	"restaurant/core/restaurant/service"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func InitTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	tracer, closer, err := cfg.New(
		serviceName,
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		log.Fatalf("ERROR: cannot init Jaeger: %v\n", err)
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

type tracingMiddleware struct {
	tracer opentracing.Tracer
	next   service.Service
}

func NewTracingMiddleware(tracer opentracing.Tracer, s service.Service) service.Service {
	return &tracingMiddleware{tracer, s}
}

func (mw tracingMiddleware) Create(restaurant domain.Restaurant) (string, error) {
	span := mw.tracer.StartSpan("Create")
	defer span.Finish()
	return mw.next.Create(restaurant)
}

func (mw tracingMiddleware) Get(id string) (domain.Restaurant, error) {
	span := mw.tracer.StartSpan("Get")
	defer span.Finish()
	return mw.next.Get(id)
}

func (mw tracingMiddleware) Update(id string, restaurant domain.Restaurant) error {
	span := mw.tracer.StartSpan("Update")
	defer span.Finish()
	return mw.next.Update(id, restaurant)
}

func (mw tracingMiddleware) Delete(id string) error {
	span := mw.tracer.StartSpan("Delete")
	defer span.Finish()
	return mw.next.Delete(id)
}

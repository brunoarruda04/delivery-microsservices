package tracing

import (
	"authentication/internal/domain"
	"authentication/internal/service"
	"io"
	"log"

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

	tracer, closer, err := cfg.NewTracer(
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
	next   service.AuthService
}

func NewTracingMiddleware(tracer opentracing.Tracer, s service.AuthService) service.AuthService {
	return &tracingMiddleware{tracer, s}
}

func (mw tracingMiddleware) Login(username, password string) (string, error) {
	span := mw.tracer.StartSpan("Login")
	defer span.Finish()
	return mw.next.Login(username, password)
}

func (mw tracingMiddleware) Register(username, password string, role domain.Role) (string, error) {
	span := mw.tracer.StartSpan("Register")
	defer span.Finish()
	return mw.next.Register(username, password, role)
}

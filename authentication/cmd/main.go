package main

import (
	"authentication/config"
	"authentication/internal/metrics"
	"authentication/internal/middleware"
	"authentication/internal/repository"
	"authentication/internal/service"
	"authentication/internal/tracing"
	"authentication/internal/transport"
	"authentication/proto"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/log"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func main() {
	// Set up logger
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// Load Environment Variables
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log("error", err)
	}

	// Set up Tracing
	tracer, closer := tracing.InitTracer("authentication-service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// Set up Database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Repository start
	repo := repository.NewSQLCRepository(db)

	// Service configurations
	svc := service.NewAuthService(repo)
	svc = middleware.NewLoggingMiddleware(logger, svc)
	svc = tracing.NewTracingMiddleware(tracer, svc)
	svc = metrics.NewMetricsMiddleware(svc)

	// Configurate the gRPC Server
	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, transport.NewGRPCServer(svc))

	// Start gRPC Server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
		if err != nil {
			panic(err)
		}

		logger.Log("transport", "HTTP", "addr", cfg.GRPCPort)

		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	// Start Metrics Server
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("transport", "HTTP", "addr", "9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

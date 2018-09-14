package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/disruptive-technologies/micro/auth"
	"bitbucket.org/disruptive-technologies/micro/health"
	"bitbucket.org/disruptive-technologies/micro/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Service implements a registration interface for services to attach
// themselves to the grpc.Server.
type Service interface {
	Register(*grpc.Server)
}

// ListenAndServe opens a tcp listener used by the grpc.Server, and registers
// each Service with the grpc.Server.
func ListenAndServe(cfg *Config, services ...Service) error {
	return ListenAndServeContext(context.Background(), cfg, services...)
}

// ListenAndServeContext opens a tcp listener used by the grpc.Server, and registers
// each Service with the grpc.Server. If the context is canceled or times out,
// the GRPC server will atempt a graceful shutdown.
func ListenAndServeContext(ctx context.Context, cfg *Config, services ...Service) error {
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	var (
		opts               = cfg.ServerOptions
		streamInterceptors = []grpc.StreamServerInterceptor{grpc_prometheus.StreamServerInterceptor}
		unaryInterceptors  = []grpc.UnaryServerInterceptor{grpc_prometheus.UnaryServerInterceptor}
	)
	if cfg.Auth != nil {
		streamInterceptors = append(streamInterceptors, auth.StreamInterceptor(cfg.Auth))
		unaryInterceptors = append(unaryInterceptors, auth.UnaryInterceptor(cfg.Auth))
	}

	// Configure tracing
	if shouldTrace, err := tracing.ConfigureTracing(ctx, cfg.TraceConfig); err != nil {
		return err
	} else if shouldTrace {
		unaryInterceptors = append(unaryInterceptors, tracing.UnaryServerInterceptor())
		streamInterceptors = append(streamInterceptors, tracing.StreamServerInterceptor())
	}

	if len(streamInterceptors) > 0 {
		opts = append(
			opts,
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
		)
	}

	if len(unaryInterceptors) > 0 {
		opts = append(
			opts,
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		)
	}

	if cfg.TLSCertFile != "" && cfg.TLSKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(cfg.TLSCertFile, cfg.TLSKeyFile)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}
	gSrv := grpc.NewServer(opts...)
	for _, s := range services {
		s.Register(gSrv)
	}
	// Make sure Prometheus metrics are initialized.
	grpc_prometheus.Register(gSrv)

	// Attach HTTP handlers
	http.Handle(cfg.ReadinessPath, health.Readiness())
	http.Handle(cfg.LivenessPath, health.Liveness(cfg.HealthChecks...))
	http.Handle(cfg.MetricsPath, prometheus.Handler())

	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)

	// Start the GRPC server
	go func(srv *grpc.Server, l net.Listener, ec chan error) {
		ec <- srv.Serve(l)
	}(gSrv, lis, errChan)

	// Serve the internal http server for metrics and health checks.
	go func(ec chan error) {
		if err := http.ListenAndServe(cfg.InternalAddress, nil); err != nil {
			ec <- err
		}
	}(errChan)

	for {
		select {
		case <-ctx.Done():
			gSrv.GracefulStop()
			return ctx.Err()
		case err := <-errChan:
			return err
		case s := <-sigChan:
			switch s {
			case os.Interrupt, syscall.SIGTERM:
				gSrv.GracefulStop()
			case os.Kill, syscall.SIGKILL:
				gSrv.Stop()
			}
			// We're still waiting for srv.Serve to return to errChan,
			// so we don't explicitly call return here yet.
		}
	}
}

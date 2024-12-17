package grpc_app

import (
	"context"
	"fmt"
	"github.com/Daniel20021510/sso/internal/grpc/auth"
	"github.com/Daniel20021510/sso/internal/grpc/interseptor"
	"github.com/Daniel20021510/sso/internal/service/auth"
	"github.com/Daniel20021510/sso/pkg/logger"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server app.
func New(authService *auth.Service, port int) *App {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// TODO: recover, metrics, trace, masking in log, auth
			interseptor.Logger,
		),
	)

	auth_grpc.RegisterServer(grpcServer, authService)

	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	logger.Infow(context.Background(), "starting grpc server", "addr", l.Addr().String())

	if err = a.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	logger.Infow(context.Background(), "stopping gRPC server", "port", a.port)

	a.gRPCServer.GracefulStop()
}

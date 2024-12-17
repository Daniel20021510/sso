package main

import (
	"context"
	"github.com/Daniel20021510/sso/internal/app"
	"github.com/Daniel20021510/sso/internal/config"
	"github.com/Daniel20021510/sso/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	defer log.Sync()

	application := app.New(cfg.GRPC.Port, cfg.PostgresConnString, cfg.TokenTTL)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	application.Postgres.Disconnect()
	logger.Infow(ctx, "Gracefully stopped")
}

func setupLogger(env string) *logger.Logger {
	var zapConf zap.Config

	var log *logger.Logger

	switch env {
	case envLocal:
		zapConf = zap.NewDevelopmentConfig()
	case envDev:
		zapConf = zap.NewDevelopmentConfig()
	case envProd:
		zapConf = zap.NewProductionConfig()
	}

	zapConf.ErrorOutputPaths = []string{"stdout"}

	log = logger.NewLogger(zapConf)

	return log
}

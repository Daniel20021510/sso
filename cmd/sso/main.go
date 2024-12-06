package main

import (
	"context"
	"github.com/Daniel20021510/sso/internal/config"
	"github.com/Daniel20021510/sso/pkg/logger"
	"go.uber.org/zap"
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

	logger.Infow(ctx, "Application started")

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
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

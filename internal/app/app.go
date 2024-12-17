package app

import (
	"github.com/Daniel20021510/sso/internal/app/grpc"
	postgres_app "github.com/Daniel20021510/sso/internal/app/postgres"
	"github.com/Daniel20021510/sso/internal/repository/postgres/app"
	"github.com/Daniel20021510/sso/internal/repository/postgres/user"
	"github.com/Daniel20021510/sso/internal/service/auth"
	"time"
)

type App struct {
	GRPCServer *grpc_app.App
	Postgres   *postgres_app.App
}

func New(grpcPort int, postgresConnString string, tokenTTL time.Duration) *App {
	postgresApp := postgres_app.New(postgresConnString)

	userRepo := user_repository.New(postgresApp.Conn())
	appRepo := app_repository.New(postgresApp.Conn())

	authService := auth.NewService(userRepo, appRepo, tokenTTL)

	grpcApp := grpc_app.New(authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

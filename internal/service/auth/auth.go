package auth

import (
	"context"
	"github.com/Daniel20021510/sso/internal/domain/model"
	"time"
)

type UserRepository interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uint64, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type AppRepository interface {
	FindByID(ctx context.Context, id uint32) (*model.App, error)
}

type Service struct {
	userRepository UserRepository
	appRepository  AppRepository
	tokenTTL       time.Duration
}

func NewService(userRepository UserRepository, appRepository AppRepository, tokenTTL time.Duration) *Service {
	return &Service{
		userRepository: userRepository,
		appRepository:  appRepository,
		tokenTTL:       tokenTTL,
	}
}

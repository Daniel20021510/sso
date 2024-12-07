package auth

import (
	"context"
	"errors"
	"github.com/Daniel20021510/sso/internal/lib/jwt"
	"github.com/Daniel20021510/sso/internal/repository"
	"github.com/Daniel20021510/sso/internal/service"
	"github.com/Daniel20021510/sso/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Login checks if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (s *Service) Login(ctx context.Context, email string, password string, appID uint32) (string, error) {
	logger.Debugw(ctx, "attempting to login user", "email", email, "app_id", appID)

	user, err := s.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Errorw(ctx, "user not found", "email", email, "error", err)

			return "", service.ErrInvalidCredentials
		}

		logger.Errorw(ctx, "failed to find user", "email", email, "error", err)

		return "", err
	}

	if err = bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		logger.Infow(ctx, "invalid credentials", "email", email, "error", err)

		return "", service.ErrInvalidCredentials
	}

	app, err := s.appRepository.FindAppByID(ctx, appID)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			logger.Errorw(ctx, "app not found", "email", email, "app_id", appID, "error", err)

			return "", service.ErrInvalidCredentials
		}

		logger.Errorw(ctx, "failed to find app", "email", email, "app_id", appID, "error", err)

		return "", err
	}

	logger.Infow(ctx, "user logged in successfully", "email", email, "app_id", appID)

	token, err := jwt.NewToken(user, app, s.tokenTTL)
	if err != nil {
		logger.Errorw(ctx, "failed to generate token", "email", email, "error", err)

		return "", err
	}

	return token, nil
}

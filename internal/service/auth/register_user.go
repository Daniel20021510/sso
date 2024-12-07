package auth

import (
	"context"
	"github.com/Daniel20021510/sso/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (s *Service) RegisterUser(ctx context.Context, email string, pass string) (uint64, error) {
	logger.Debugw(ctx, "registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorw(ctx, "failed to hash password", "error", err)

		return 0, err
	}

	id, err := s.userRepository.SaveUser(ctx, email, passHash)
	if err != nil {
		logger.Errorw(ctx, "failed to save user", "error", err)

		return 0, err
	}

	return id, nil
}

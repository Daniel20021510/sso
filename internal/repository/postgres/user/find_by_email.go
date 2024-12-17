package user_repository

import (
	"context"
	"errors"
	"github.com/Daniel20021510/sso/internal/domain/model"
	"github.com/Daniel20021510/sso/internal/repository"
	"github.com/jackc/pgx/v5"
)

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := ur.q.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		return nil, err
	}

	return &model.User{
		ID:       uint64(user.ID),
		Email:    user.Email,
		PassHash: []byte(user.PassHash),
	}, nil
}

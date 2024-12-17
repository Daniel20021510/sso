package user_repository

import (
	"context"
	"errors"
	"github.com/Daniel20021510/sso/internal/repository"
	"github.com/Daniel20021510/sso/internal/repository/postgres"
	"github.com/Daniel20021510/sso/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

func (ur *UserRepository) SaveUser(ctx context.Context, email string, passHash []byte) (uint64, error) {
	id, err := ur.q.CreateUser(ctx, &sqlc_postgres.CreateUserParams{Email: email, PassHash: string(passHash)})
	if err != nil {
		var pgxErr pgconn.PgError
		if errors.As(err, &pgxErr) && pgxErr.Code == postgres.UniqueConstraintViolation {
			return 0, repository.ErrUserExists
		}

		return 0, err
	}

	return uint64(id), nil
}

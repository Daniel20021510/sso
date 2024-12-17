package user_repository

import (
	"github.com/Daniel20021510/sso/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	q    sqlc_postgres.Querier
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		q:    sqlc_postgres.New(conn),
		conn: conn,
	}
}

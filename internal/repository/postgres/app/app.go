package app_repository

import (
	"github.com/Daniel20021510/sso/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

type AppRepository struct {
	q    sqlc_postgres.Querier
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *AppRepository {
	return &AppRepository{
		q:    sqlc_postgres.New(conn),
		conn: conn,
	}
}

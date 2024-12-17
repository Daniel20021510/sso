package postgres_app

import (
	"context"
	"fmt"
	"github.com/Daniel20021510/sso/pkg/logger"
	"github.com/jackc/pgx/v5"
)

type App struct {
	connString string
	conn       *pgx.Conn
}

// New creates new postgres app.
func New(connString string) *App {
	return &App{
		connString: connString,
	}
}

// MustConnect connects to the database and panics if any error occurs.
func (a *App) MustConnect() {
	if err := a.Connect(); err != nil {
		panic(err)
	}
}

// Connect connects to the database.
func (a *App) Connect() error {
	logger.Infow(context.Background(), "connecting to database")
	dbConn, err := pgx.Connect(context.Background(), a.connString)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	a.conn = dbConn

	return nil
}

// Conn return the underlying *pgx.Conn.
func (a *App) Conn() *pgx.Conn {
	return a.conn
}

// Disconnect disconnects from the database.
func (a *App) Disconnect() {
	logger.Infow(context.Background(), "disconnecting from database")

	a.conn.Close(context.Background())
}

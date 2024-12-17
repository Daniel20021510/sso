package app_repository

import (
	"context"
	"errors"
	"github.com/Daniel20021510/sso/internal/domain/model"
	"github.com/Daniel20021510/sso/internal/repository"
	"github.com/jackc/pgx/v5"
)

func (ar *AppRepository) FindByID(ctx context.Context, id uint32) (*model.App, error) {
	app, err := ar.q.GetApp(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrAppNotFound
		}

		return nil, err
	}

	return &model.App{
		ID:        uint32(app.ID),
		Name:      app.Name,
		SecretKey: app.Secret,
	}, nil
}

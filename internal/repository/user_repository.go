package repository

import (
	"context"
	"wallet-service/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetUserForUpdate(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateBalance(ctx context.Context, tx pgx.Tx, id uuid.UUID, newBalance float64) error
}

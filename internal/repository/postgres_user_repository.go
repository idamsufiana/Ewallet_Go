package repository

import (
	"context"
	"errors"
	"wallet-service/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

const (
	queryGetUserForUpdate = `
		SELECT id, name, balance
		FROM users
		WHERE id=$1
		FOR UPDATE
	`

	queryGetUser = `
		SELECT id, name, balance
		FROM users
		WHERE id=$1
	`

	queryUpdateBalance = `
		UPDATE users
		SET balance=$1
		WHERE id=$2
	`
)

func (r *PostgresUserRepository) GetUserForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
) (*model.User, error) {

	var user model.User

	err := tx.QueryRow(ctx, queryGetUserForUpdate, id).
		Scan(&user.ID, &user.Name, &user.Balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUser(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
) (*model.User, error) {

	var user model.User

	err := tx.QueryRow(ctx, queryGetUser, id).
		Scan(&user.ID, &user.Name, &user.Balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUserByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, queryGetUser, id).
		Scan(&user.ID, &user.Name, &user.Balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) UpdateBalance(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	newBalance float64,
) error {

	cmdTag, err := tx.Exec(ctx, queryUpdateBalance, newBalance, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

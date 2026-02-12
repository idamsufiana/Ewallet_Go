package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func NewPostgres() (*pgxpool.Pool, error) {
	dsn := viper.GetString("database.url")

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

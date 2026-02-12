package service

import (
	"context"
	"errors"
	"wallet-service/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletService struct {
	db   *pgxpool.Pool
	repo repository.UserRepository
}

// Constructor
func NewWalletService(db *pgxpool.Pool, repo repository.UserRepository) *WalletService {
	return &WalletService{
		db:   db,
		repo: repo,
	}
}

// Withdraw mengurangi saldo user secara atomik
func (s *WalletService) Withdraw(ctx context.Context, userID uuid.UUID, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	// Mulai transaksi
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx) // rollback otomatis jika commit gagal

	// Ambil user dengan row lock
	user, err := s.repo.GetUserForUpdate(ctx, tx, userID)
	if err != nil {
		return 0, err
	}

	if user.Balance < amount {
		return 0, errors.New("insufficient balance")
	}

	newBalance := user.Balance - amount

	// Update saldo
	if err := s.repo.UpdateBalance(ctx, tx, userID, newBalance); err != nil {
		return 0, err
	}

	// Commit transaksi
	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return newBalance, nil
}

// GetBalance cukup baca tanpa transaksi
func (s *WalletService) GetBalance(ctx context.Context, userID uuid.UUID) (float64, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return user.Balance, nil
}

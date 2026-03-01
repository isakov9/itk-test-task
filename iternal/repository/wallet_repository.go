package repository

import (
	"context"
	"errors"
)

var ErrWalletNotFound = errors.New("wallet not found")

type WalletRepository interface {
	GetWalletById(ctx context.Context, walletID string) (int64, error)
	UpdateBalance(ctx context.Context, walletID string, amount int64) error
}

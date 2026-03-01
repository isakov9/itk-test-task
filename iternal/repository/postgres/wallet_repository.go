package postgres

import (
	"context"
	"errors"
	"fmt"
	"itk-test-task/iternal/repository"

	"github.com/jackc/pgx/v5"
)

type WalletRepository struct {
	Repository
}

func NewWalletRepository(connector *DBConnector) repository.WalletRepository {
	return &WalletRepository{Repository{pool: connector.Pool}}
}

func (r Repository) GetWalletById(ctx context.Context, walletID string) (int64, error) {
	query := `
        SELECT amount
        FROM itk.wallets
        WHERE wallet_id = $1;
    `

	var amount int64

	err := r.pool.QueryRow(ctx, query, walletID).Scan(&amount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, repository.ErrWalletNotFound
		}
		return 0, fmt.Errorf("failed to get wallet: %w", err)
	}

	return amount, nil
}

func (r Repository) UpdateBalance(ctx context.Context, walletID string, amount int64) error {
	query := `
		INSERT INTO itk.wallets (wallet_id, amount)
		VALUES ($1, $2)
		ON CONFLICT (wallet_id)
		DO UPDATE SET amount = itk.wallets.amount + EXCLUDED.amount
		RETURNING amount;
	`

	var newAmount int64
	err := r.pool.QueryRow(ctx, query, walletID, amount).Scan(&newAmount)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	return nil
}

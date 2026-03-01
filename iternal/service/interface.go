package service

import "context"

type WalletServiceInterface interface {
	UpdateWallet(ctx context.Context, input *UpdateWallet) error
	GetWalletById(ctx context.Context, walletID string) (int64, error)
}

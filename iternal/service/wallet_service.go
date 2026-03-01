package service

import (
	"context"
	"itk-test-task/iternal/repository"
)

type WalletService struct {
	WalletRepository repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) *WalletService {
	return &WalletService{
		WalletRepository: repo,
	}
}

func (s *WalletService) UpdateWallet(ctx context.Context, input *UpdateWallet) error {
	if input.OperationType == "WITHDRAW" {
		input.Amount *= -1
	}

	err := s.WalletRepository.UpdateBalance(ctx, input.WalletID, input.Amount)

	if err != nil {
		return err
	}

	return nil
}

func (s *WalletService) GetWalletById(ctx context.Context, walletID string) (int64, error) {
	amount, err := s.WalletRepository.GetWalletById(ctx, walletID)

	if err != nil {
		return 0, err
	}

	return amount, nil
}

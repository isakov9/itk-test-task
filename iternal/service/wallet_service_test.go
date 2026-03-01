package service_test

import (
	"context"
	"testing"

	"itk-test-task/iternal/mocks"
	"itk-test-task/iternal/repository"
	"itk-test-task/iternal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateWallet_Deposit(t *testing.T) {
	mockRepo := mocks.NewWalletRepository(t)
	mockRepo.On("UpdateBalance", mock.Anything, "some-uuid", int64(1000)).Return(nil)

	svc := service.NewWalletService(mockRepo)

	err := svc.UpdateWallet(context.Background(), &service.UpdateWallet{
		WalletID:      "some-uuid",
		OperationType: service.OperationDeposit,
		Amount:        1000,
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWallet_Withdraw(t *testing.T) {
	mockRepo := mocks.NewWalletRepository(t)
	mockRepo.On("UpdateBalance", mock.Anything, "some-uuid", int64(-500)).Return(nil)

	svc := service.NewWalletService(mockRepo)

	err := svc.UpdateWallet(context.Background(), &service.UpdateWallet{
		WalletID:      "some-uuid",
		OperationType: service.OperationWithdraw,
		Amount:        500,
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetWalletById_NotFound(t *testing.T) {
	mockRepo := mocks.NewWalletRepository(t)
	mockRepo.On("GetWalletById", mock.Anything, "some-uuid").Return(int64(0), repository.ErrWalletNotFound)

	svc := service.NewWalletService(mockRepo)

	_, err := svc.GetWalletById(context.Background(), "some-uuid")

	assert.ErrorIs(t, err, repository.ErrWalletNotFound)
}

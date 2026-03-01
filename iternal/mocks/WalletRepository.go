package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

type WalletRepository struct {
	mock.Mock
}

func (_m *WalletRepository) GetWalletById(ctx context.Context, walletID string) (int64, error) {
	ret := _m.Called(ctx, walletID)

	if len(ret) == 0 {
		panic("no return value specified for GetWalletById")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, walletID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, walletID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, walletID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *WalletRepository) UpdateBalance(ctx context.Context, walletID string, amount int64) error {
	ret := _m.Called(ctx, walletID, amount)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, walletID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func NewWalletRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

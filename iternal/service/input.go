package service

const (
	OperationDeposit  = "DEPOSIT"
	OperationWithdraw = "WITHDRAW"
)

type UpdateWallet struct {
	WalletID      string
	OperationType string
	Amount        int64
}

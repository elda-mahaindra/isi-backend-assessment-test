package service

import (
	"context"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/dto"

	"github.com/sirupsen/logrus"
)

type IService interface {
	BalanceCheck(ctx context.Context, request dto.BalanceCheckRequest) (saldo int64, err error)
	Deposit(ctx context.Context, request dto.DepositRequest) (saldo int64, err error)
	Registration(ctx context.Context, request dto.RegistrationRequest) (noRekening string, err error)
	TransactionHistory(ctx context.Context, request dto.TransactionHistoryRequest) (entries []sqlc.Entry, err error)
	Withdrawal(ctx context.Context, request dto.WithdrawalRequest) (saldo int64, err error)
}

type Service struct {
	logger *logrus.Logger
	store  db.IStore
}

func NewService(logger *logrus.Logger, store db.IStore) IService {
	return &Service{
		logger: logger,
		store:  store,
	}
}

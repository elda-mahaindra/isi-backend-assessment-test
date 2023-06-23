package db

import (
	"context"

	"tabungan-api/db/sqlc"
)

type IStore interface {
	sqlc.Querier

	DepositTx(ctx context.Context, arg DepositTxParams) (DepositTxResult, error)
	WithdrawalTx(ctx context.Context, arg WithdrawalTxParams) (WithdrawalTxResult, error)
}

type DepositTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type DepositTxResult struct {
	Saldo int64 `json:"saldo"`
}

type WithdrawalTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type WithdrawalTxResult struct {
	Saldo int64 `json:"saldo"`
}

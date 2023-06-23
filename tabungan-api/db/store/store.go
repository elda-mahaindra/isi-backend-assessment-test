package db

import (
	"context"

	"tabungan-api/db/sqlc"
)

type IStore interface {
	sqlc.Querier

	DepositTx(ctx context.Context, arg DepositTxParams) (DepositTxResult, error)
	RegistrationTx(ctx context.Context, arg RegistrationTxParams) (RegistrationTxResult, error)
	WithdrawalTx(ctx context.Context, arg WithdrawalTxParams) (WithdrawalTxResult, error)
}

type DepositTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type DepositTxResult struct {
	Saldo int64 `json:"saldo"`
}

type RegistrationTxParams struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type RegistrationTxResult struct {
	NoRekening string `json:"no_rekening"`
}

type WithdrawalTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type WithdrawalTxResult struct {
	Saldo int64 `json:"saldo"`
}

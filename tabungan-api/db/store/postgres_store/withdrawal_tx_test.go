package postgres_store

import (
	"context"
	"testing"

	db "tabungan-api/db/store"
	"tabungan-api/utils/random"

	"github.com/stretchr/testify/require"
)

func withdrawRandomAmount(t *testing.T, noRekening string) db.WithdrawalTxResult {
	store := NewPostgresStore(logger, testDB)

	// get saldo before withdrawal
	saldo := func(noRekening string) int64 {
		account, err := store.GetAccount(context.Background(), noRekening)
		require.NoError(t, err)

		return account.Saldo
	}(noRekening)

	arg := db.WithdrawalTxParams{
		Nominal:    int64(random.GenerateNumber(50000, 100000)),
		NoRekening: noRekening,
	}

	result, err := store.WithdrawalTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check entry
	entry := result.Entry
	require.Equal(t, "D", entry.Code)
	require.Equal(t, arg.Nominal, entry.Nominal)
	require.Equal(t, arg.NoRekening, entry.NoRekening)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	// check account
	account := result.Account
	require.Equal(t, arg.NoRekening, account.NoRekening)
	require.Equal(t, saldo-arg.Nominal, account.Saldo)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.NoRekening)

	return result
}

func TestWithdrawalTx(t *testing.T) {
	result := registerRandomCustomer(t)

	account := result.Account

	depositRandomAmount(t, account.NoRekening)
	withdrawRandomAmount(t, account.NoRekening)
}

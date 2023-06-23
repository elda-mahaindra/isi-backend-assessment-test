package postgres_store

import (
	"context"
	"testing"

	db "tabungan-api/db/store"
	"tabungan-api/utils/random"

	"github.com/stretchr/testify/require"
)

func depositRandomAmount(t *testing.T, noRekening string) db.DepositTxResult {
	store := NewPostgresStore(logger, testDB)

	// get saldo before deposit
	saldo := func(noRekening string) int64 {
		account, err := store.GetAccount(context.Background(), noRekening)
		require.NoError(t, err)

		return account.Saldo
	}(noRekening)

	arg := db.DepositTxParams{
		Nominal:    int64(random.GenerateNumber(200000, 400000)),
		NoRekening: noRekening,
	}

	result, err := store.DepositTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check entry
	entry := result.Entry
	require.Equal(t, "C", entry.Code)
	require.Equal(t, arg.Nominal, entry.Nominal)
	require.Equal(t, arg.NoRekening, entry.NoRekening)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	// check account
	account := result.Account
	require.Equal(t, arg.NoRekening, account.NoRekening)
	require.Equal(t, saldo+arg.Nominal, account.Saldo)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.NoRekening)

	return result
}

func TestDepositTx(t *testing.T) {
	result := registerRandomCustomer(t)

	account := result.Account

	depositRandomAmount(t, account.NoRekening)
}

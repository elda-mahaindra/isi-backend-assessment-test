package postgres_store

import (
	"context"
	"testing"

	db "tabungan-api/db/store"
	"tabungan-api/utils/random"

	"github.com/stretchr/testify/require"
)

func registerRandomCustomer(t *testing.T) db.RegistrationTxResult {
	store := NewPostgresStore(logger, testDB)

	arg := db.RegistrationTxParams{
		Nama: random.GenerateAlphabetString(12),
		Nik:  random.GenerateNumericString(16),
		NoHp: random.GeneratePhoneNo(),
	}

	result, err := store.RegistrationTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check customer
	customer := result.Customer
	require.Equal(t, arg.Nama, customer.Nama)
	require.Equal(t, arg.Nik, customer.Nik)
	require.Equal(t, arg.NoHp, customer.NoHp)

	require.NotZero(t, customer.CreatedAt)
	require.NotZero(t, customer.ID)

	// check account
	account := result.Account
	require.Equal(t, customer.ID, account.CustomerID)
	require.Equal(t, int64(0), account.Saldo)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.NoRekening)

	return result
}

func TestRegistrationTx(t *testing.T) {
	registerRandomCustomer(t)
}

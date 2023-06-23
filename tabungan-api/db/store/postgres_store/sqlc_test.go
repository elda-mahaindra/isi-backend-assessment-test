package postgres_store

import (
	"context"
	"testing"

	"tabungan-api/db/sqlc"
	"tabungan-api/utils/random"

	"github.com/stretchr/testify/require"
)

func TestGetEntries(t *testing.T) {
	result := registerRandomCustomer(t)

	account := result.Account

	// create random entries
	entriesNumber := random.GenerateNumber(1, 10)

	entriesCreated := []sqlc.Entry{}

	for i := 0; i < entriesNumber; i++ {
		if i%2 == 0 {
			result := depositRandomAmount(t, account.NoRekening)

			entriesCreated = append(entriesCreated, result.Entry)
		} else {
			result := withdrawRandomAmount(t, account.NoRekening)

			entriesCreated = append(entriesCreated, result.Entry)
		}
	}

	entries, err := testQueries.GetEntries(context.Background(), account.NoRekening)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, entriesNumber, len(entries))

	for index, entry := range entries {
		require.Equal(t, entriesCreated[index], entry)
	}
}

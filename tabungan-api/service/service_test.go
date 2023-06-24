package service_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/db/store/postgres_store"
	"tabungan-api/dto"
	"tabungan-api/service"
	"tabungan-api/utils/config"
	"tabungan-api/utils/random"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var logger *logrus.Logger
var testStore db.IStore

func init() {
	logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)                     //default
	logger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	// load environment variables from .env file
	config, err := config.LoadConfig("../.")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %s", err)
	}

	testStore = postgres_store.NewPostgresStore(logger, conn)
}

func createService() service.IService {
	service := service.NewService(logger, testStore)

	return service
}

func registerRandomCustomer(t *testing.T) db.RegistrationTxResult {
	result, err := testStore.RegistrationTx(context.Background(), db.RegistrationTxParams{
		Nama: random.GenerateAlphabetString(12),
		Nik:  random.GenerateNumericString(16),
		NoHp: random.GeneratePhoneNo(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	return result
}

func depositRandomAmount(t *testing.T, noRekening string) db.DepositTxResult {
	result, err := testStore.DepositTx(context.Background(), db.DepositTxParams{
		Nominal:    int64(random.GenerateNumber(200000, 400000)),
		NoRekening: noRekening,
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)

	return result
}

func withdrawRandomAmount(t *testing.T, noRekening string) db.WithdrawalTxResult {
	result, err := testStore.WithdrawalTx(context.Background(), db.WithdrawalTxParams{
		Nominal:    int64(random.GenerateNumber(50000, 100000)),
		NoRekening: noRekening,
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)

	return result
}

func createRandomEntries(t *testing.T, entriesNumber int, noRekening string) []sqlc.Entry {
	// create random entries
	entriesCreated := []sqlc.Entry{}

	for i := 0; i < entriesNumber; i++ {
		if i%2 == 0 {
			result := depositRandomAmount(t, noRekening)

			entriesCreated = append(entriesCreated, result.Entry)
		} else {
			result := withdrawRandomAmount(t, noRekening)

			entriesCreated = append(entriesCreated, result.Entry)
		}
	}

	return entriesCreated
}

func TestBalanceCheck(t *testing.T) {
	t.Parallel()

	createdAccount := registerRandomCustomer(t).Account

	// TDD test
	testCases := []struct {
		name    string
		request dto.BalanceCheckRequest
	}{
		{
			name: "ok",
			request: dto.BalanceCheckRequest{
				NoRekening: createdAccount.NoRekening,
			},
		},
		{
			name: "not ok - account not found",
			request: dto.BalanceCheckRequest{
				NoRekening: "wrong no rekening",
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			service := createService()

			saldo, err := service.BalanceCheck(context.Background(), testCase.request)

			if testCase.name == "ok" {
				require.NoError(t, err)
				require.Zero(t, saldo)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	t.Parallel()

	createdAccount := registerRandomCustomer(t).Account

	// TDD test
	testCases := []struct {
		name    string
		request dto.DepositRequest
	}{
		{
			name: "ok",
			request: dto.DepositRequest{
				Nominal:    int64(random.GenerateNumber(200000, 400000)),
				NoRekening: createdAccount.NoRekening,
			},
		},
		{
			name: "not ok - account not found",
			request: dto.DepositRequest{
				Nominal:    int64(random.GenerateNumber(200000, 400000)),
				NoRekening: "wrong no rekening",
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			service := createService()

			saldo, err := service.Deposit(context.Background(), testCase.request)

			if testCase.name == "ok" {
				require.NoError(t, err)
				require.NotEmpty(t, saldo)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestRegistration(t *testing.T) {
	t.Parallel()

	// register customer first so we can try to register another customer with the same nik or no hp
	// we'll do it directly from data access layer
	createdCustomer := registerRandomCustomer(t).Customer

	// TDD test
	testCases := []struct {
		name    string
		request dto.RegistrationRequest
	}{
		{
			name: "ok",
			request: dto.RegistrationRequest{
				Nama: random.GenerateAlphabetString(12),
				Nik:  random.GenerateNumericString(16),
				NoHp: random.GeneratePhoneNo(),
			},
		},
		{
			name: "not ok - nik duplication",
			request: dto.RegistrationRequest{
				Nama: random.GenerateAlphabetString(12),
				Nik:  createdCustomer.Nik,
				NoHp: random.GeneratePhoneNo(),
			},
		},
		{
			name: "not ok - no hp duplication",
			request: dto.RegistrationRequest{
				Nama: random.GenerateAlphabetString(12),
				Nik:  random.GenerateNumericString(16),
				NoHp: createdCustomer.NoHp,
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			service := createService()

			noRekening, err := service.Registration(context.Background(), testCase.request)

			if testCase.name == "ok" {
				require.NoError(t, err)
				require.NotEmpty(t, noRekening)
			} else {
				require.Error(t, err)
				require.Zero(t, noRekening)
			}
		})
	}
}

func TestTransactionHistory(t *testing.T) {
	t.Parallel()

	// TDD test
	testCases := []struct {
		entriesNumber int
		name          string
	}{
		{
			entriesNumber: random.GenerateNumber(1, 10),
			name:          "ok",
		},
		{
			entriesNumber: random.GenerateNumber(1, 10),
			name:          "not ok - account not found",
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			service := createService()

			createdAccount := registerRandomCustomer(t).Account

			entriesCreated := createRandomEntries(t, testCase.entriesNumber, createdAccount.NoRekening)

			noRekening := "wrong no rekening"

			if testCase.name != "not ok - account not found" {
				noRekening = createdAccount.NoRekening
			}

			entries, err := service.TransactionHistory(context.Background(), dto.TransactionHistoryRequest{
				NoRekening: noRekening,
			})

			if testCase.name == "ok" {
				require.NoError(t, err)
				require.NotEmpty(t, entries)

				require.Equal(t, len(entriesCreated), len(entries))
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestWithdrawal(t *testing.T) {
	t.Parallel()

	// TDD test
	testCases := []struct {
		name              string
		nominalDeposit    int64
		nominalWithdrawal int64
	}{
		{
			name:              "ok",
			nominalDeposit:    int64(random.GenerateNumber(200000, 400000)),
			nominalWithdrawal: int64(random.GenerateNumber(50000, 100000)),
		},
		{
			name:              "not ok - account not found",
			nominalDeposit:    int64(random.GenerateNumber(200000, 400000)),
			nominalWithdrawal: int64(random.GenerateNumber(50000, 100000)),
		},
		{
			name:              "not ok - insufficient balance",
			nominalDeposit:    int64(random.GenerateNumber(200000, 400000)),
			nominalWithdrawal: int64(random.GenerateNumber(500000, 1000000)),
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			service := createService()

			createdAccount := registerRandomCustomer(t).Account

			_, err := service.Deposit(context.Background(), dto.DepositRequest{
				Nominal:    testCase.nominalDeposit,
				NoRekening: createdAccount.NoRekening,
			})
			require.NoError(t, err)

			noRekening := "wrong no rekening"

			if testCase.name != "not ok - account not found" {
				noRekening = createdAccount.NoRekening
			}

			saldo, err := service.Withdrawal(context.Background(), dto.WithdrawalRequest{
				Nominal:    testCase.nominalWithdrawal,
				NoRekening: noRekening,
			})

			if testCase.name == "ok" {
				require.NoError(t, err)
				require.NotEmpty(t, saldo)
			} else {
				require.Error(t, err)
			}
		})
	}
}

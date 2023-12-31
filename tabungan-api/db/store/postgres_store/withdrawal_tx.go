package postgres_store

import (
	"context"
	"fmt"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (store *PostgresStore) WithdrawalTx(ctx context.Context, arg db.WithdrawalTxParams) (db.WithdrawalTxResult, error) {
	const op errs.Op = "postgres_store/WithdrawalTx"

	var result db.WithdrawalTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// create entry
		result.Entry, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			Code:       "D",
			Nominal:    arg.Nominal,
			NoRekening: arg.NoRekening,
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'CreateEntry' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		// get account
		account, err := q.GetAccount(ctx, arg.NoRekening)
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'GetAccount' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		// check balance left
		if account.Saldo < arg.Nominal {
			e := errs.E(op, errs.Database, fmt.Sprintf("insufficient balance: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return e
		}

		// update saldo
		result.Account, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NoRekening: arg.NoRekening,
			Saldo:      account.Saldo - arg.Nominal,
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'UpdateSaldo' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		return err
	})

	return result, err
}

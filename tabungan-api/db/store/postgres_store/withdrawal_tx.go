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
		_, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			NoRekening: arg.NoRekening,
			Nominal:    arg.Nominal,
			TypeID:     "D",
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'CreateEntry' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Trace(e.Error())

			return e
		}

		// get account
		account, err := q.GetAccount(ctx, arg.NoRekening)
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'GetAccount' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Trace(e.Error())

			return e
		}

		// update saldo
		_, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NoRekening: arg.NoRekening,
			Saldo:      account.Saldo - arg.Nominal,
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'UpdateSaldo' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Trace(e.Error())

			return e
		}

		result.Saldo = account.Saldo - arg.Nominal

		return err
	})

	return result, err
}

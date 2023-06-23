package postgres_store

import (
	"context"
	"fmt"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (store *PostgresStore) DepositTx(ctx context.Context, arg db.DepositTxParams) (db.DepositTxResult, error) {
	const op errs.Op = "postgres_store/DepositTx"

	var result db.DepositTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// create entry
		_, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			Code:       "C",
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

		// update saldo
		_, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NoRekening: arg.NoRekening,
			Saldo:      account.Saldo + arg.Nominal,
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'UpdateSaldo' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		result.Saldo = account.Saldo + arg.Nominal

		return err
	})

	return result, err
}

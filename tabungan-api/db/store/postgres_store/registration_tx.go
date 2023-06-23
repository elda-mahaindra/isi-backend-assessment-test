package postgres_store

import (
	"context"
	"fmt"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/utils/errs"
	"tabungan-api/utils/random"

	"github.com/sirupsen/logrus"
)

func (store *PostgresStore) RegistrationTx(ctx context.Context, arg db.RegistrationTxParams) (db.RegistrationTxResult, error) {
	const op errs.Op = "postgres_store/RegistrationTx"

	var result db.RegistrationTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// create customers
		result.Customer, err = q.CreateCustomer(ctx, sqlc.CreateCustomerParams{
			Nama: arg.Nama,
			Nik:  arg.Nik,
			NoHp: arg.NoHp,
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'CreateCustomer' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		// create account
		result.Account, err = q.CreateAccount(ctx, sqlc.CreateAccountParams{
			CustomerID: result.Customer.ID,
			NoRekening: random.GenerateNumericString(16),
		})
		if err != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to execute 'CreateAccount' query: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Debug(e.Error())

			return err
		}

		return err
	})

	return result, err
}

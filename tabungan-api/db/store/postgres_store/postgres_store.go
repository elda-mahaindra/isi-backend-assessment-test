package postgres_store

import (
	"context"
	"database/sql"
	"fmt"

	"tabungan-api/db/sqlc"
	db "tabungan-api/db/store"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

type PostgresStore struct {
	*sqlc.Queries
	logger *logrus.Logger
	db     *sql.DB
}

func NewPostgresStore(logger *logrus.Logger, db *sql.DB) db.IStore {
	return &PostgresStore{
		logger:  logger,
		db:      db,
		Queries: sqlc.New(db),
	}
}

// execTx executes a generic database transaction
func (store *PostgresStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	const op errs.Op = "store/PostgresStore.execTx"

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		e := errs.E(op, errs.Database, fmt.Sprintf("failed to begin tx: %s", err))

		store.logger.WithFields(logrus.Fields{
			"op": op,
		}).Trace(e.Error())

		return e
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			e := errs.E(op, errs.Database, fmt.Sprintf("failed to rollback tx: %s", err))

			store.logger.WithFields(logrus.Fields{
				"op": op,
			}).Trace(e.Error())

			return e
		}
		return err
	}

	return tx.Commit()
}

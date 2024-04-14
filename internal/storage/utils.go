package storage

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func WithTransaction(db *sql.DB, transaction func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	if err := transaction(tx); err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errors.Wrapf(err, "can't rollback with error %s", errRollback)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "can't commit transaction")
	}

	return nil
}

func checkPgConflictError(err error) error {
	var e *pgconn.PgError

	if !errors.As(err, &e) {
		return err
	}

	if e.Code == uniqueConflictCode && e.ConstraintName == uniqueConstraintName {
		return ErrorBannerConflictExists
	}

	return err
}

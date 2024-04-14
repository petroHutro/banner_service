package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (st *storageData) DeleteBanner(id uint32) error {
	var deletedID uint32

	if err := st.db.QueryRow(deleteQuery, id).
		Scan(
			&deletedID,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Wrapf(ErrorBannerNotFound, "with id %d", id)
		}

		return errors.Wrapf(err, "can't delete banner with id %d", id)
	}

	return nil
}

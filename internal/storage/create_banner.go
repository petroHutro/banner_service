package storage

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
)

func addContent(tx *sql.Tx, id uint32, content string) error {
	if _, err := tx.Exec(addContentQuery, id, content); err != nil {
		return errors.Wrap(err, "can't add content to banner")
	}

	return nil
}

func (st *storageData) CreateBanner(featureID uint32, tagIDs []uint32,
	content string, isActive bool,
) (uint32, error) {
	var createdID uint32

	if err := WithTransaction(st.db,
		func(tx *sql.Tx) error {
			if err := tx.QueryRow(createQuery, isActive).
				Scan(
					&createdID,
				); err != nil {
				return errors.Wrap(err, "can't create banner")
			}

			if err := addContent(tx, createdID, content); err != nil {
				return err
			}

			if _, err := tx.Exec(addFeaturesAndTagsQuery, createdID, featureID,
				pgtype.FlatArray[uint32](tagIDs)); err != nil {
				return errors.Wrapf(checkPgConflictError(err),
					"can't add feature id %d and tag ids %v to banner", featureID, tagIDs)
			}

			return nil
		},
	); err != nil {
		return 0, errors.Wrap(err, "when creating banner")
	}

	return createdID, nil
}

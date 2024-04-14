package storage

import (
	"banner_service/internal/models"
	"database/sql"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
)

func updateBannerInfo(tx *sql.Tx, id uint32, bnr *models.BannerUpdate) error {
	switch {
	case bnr.TagIDs != nil && bnr.FeatureID != 0:
		if _, err := tx.Exec(updateFeaturesQuery, id, bnr.FeatureID); err != nil {
			return errors.Wrapf(checkPgConflictError(err),
				"can't update feature id %d to banner", bnr.FeatureID)
		}
	case bnr.TagIDs != nil:
		var featureID uint32
		if err := tx.QueryRow(deleteFeaturesTagsQuery, id).Scan(&featureID); err != nil {
			return errors.Wrap(err, "can't delete feature id and tag ids of banner")
		}

		if bnr.FeatureID != 0 {
			featureID = bnr.FeatureID
		}

		if _, err := tx.Exec(addFeaturesAndTagsQuery, id, featureID,
			pgtype.FlatArray[uint32](bnr.TagIDs)); err != nil {
			return errors.Wrapf(checkPgConflictError(err),
				"can't add feature id %d and tag ids %v to banner", featureID, bnr.TagIDs)
		}
	}

	return nil
}

func (st *storageData) UpdateBanner(id uint32, banner *models.BannerUpdate) error {
	var updatedID uint32

	if err := WithTransaction(st.db,
		func(tx *sql.Tx) error {
			if err := tx.QueryRow(checkDeleted, id).Scan(&updatedID); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return ErrorBannerNotFound
				}

				return errors.Wrapf(err, "can't check banner on deleted")
			}

			if banner.IsActive {
				if err := tx.QueryRow(updateActiveQuery,
					id, banner.IsActive).
					Scan(&updatedID); err != nil {
					if errors.Is(err, pgx.ErrNoRows) {
						return ErrorBannerNotFound
					}

					return errors.Wrapf(err, "can't update banner")
				}
			}

			if banner.Content != nil {
				if err := addContent(tx, id, string(banner.Content)); err != nil {
					return err
				}
			}

			return updateBannerInfo(tx, id, banner)
		},
	); err != nil {
		return errors.Wrapf(err, "when updating banner with id %d", id)
	}

	return nil
}

package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (st *storageData) GetBanner(featureID, tagID uint32) (string, error) {
	var content string
	if err := st.db.QueryRow(getQuery, featureID, tagID, 1).
		Scan(
			&content,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return content, errors.Wrapf(ErrorBannerNotFound,
				"with feature id %d and tag id %d and version %v", featureID, tagID, 1)
		}

		return content, errors.Wrapf(err,
			"can't get banner with feature id %d and tag id %d and version %v", featureID, tagID, 1)
	}

	return content, nil
}

package storage

import (
	"banner_service/internal/models"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
)

func filterBanners(tx *sql.Tx, bnr *models.BannerInfo,
	offset, limit uint64,
) ([]models.Banner, error) {
	args := []any{bnr.FeatureID != 0, bnr.TagID != 0, limit, offset}
	query := filterNotNullQuery

	if bnr.FeatureID != 0 && bnr.TagID != 0 {
		args = []any{limit, offset}
		query = filterNullQuery
	}

	rows, err := tx.Query(query, args...)
	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "can't execute filter banner query")
	}

	banners := make([]models.Banner, 0)

	for rows.Next() {
		var filteredBanner models.Banner

		err := rows.Scan(
			&filteredBanner.ID,
			&filteredBanner.IsActive,
			&filteredBanner.CreatedAt,
			&filteredBanner.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "can't scan filter banner query result")
		}

		filteredBanner.TagIDs = make([]uint32, 0)
		filteredBanner.Versions = make([]models.Content, 0)

		banners = append(banners, filteredBanner)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "can't end scan filter banner query result")
	}

	return banners, nil
}

func selectTagFeatureForBanners(tx *sql.Tx, banners []models.Banner) ([]models.Banner, error) {
	bannerIDs := make([]uint32, len(banners))
	bannerIndexes := make(map[uint32]int64)

	for index, bnr := range banners {
		bannerIDs[index] = bnr.ID
		bannerIndexes[bnr.ID] = int64(index)
	}

	rows, err := tx.Query(getTagQuery, bannerIDs)
	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "can't execute get tags for banner query")
	}

	for rows.Next() {
		var tags pgtype.Array[uint32]

		var bannerID, featureID uint32

		err := rows.Scan(
			&bannerID,
			&tags,
			&featureID,
		)
		if err != nil {
			return nil, errors.Wrap(err, "can't scan get tags for banner query result")
		}

		banners[bannerIndexes[bannerID]].FeatureID = featureID

		banners[bannerIndexes[bannerID]].TagIDs = []uint32{}
		if tags.Valid {
			banners[bannerIndexes[bannerID]].TagIDs = tags.Elements
		}
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "can't end scan get tags for banner query result")
	}

	return banners, nil
}

func selectContentForBanners(tx *sql.Tx, banners []models.Banner) ([]models.Banner, error) {
	bannerIDs := make([]uint32, len(banners))
	bannerIndexes := make(map[uint32]int64)

	for index, bnr := range banners {
		bannerIDs[index] = bnr.ID
		bannerIndexes[bnr.ID] = int64(index)
	}

	rows, err := tx.Query(getVersionQuery, bannerIDs)
	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "can't execute get contents for banner query")
	}

	for rows.Next() {
		var bannerContent models.Content

		var bannerID uint32

		err := rows.Scan(
			&bannerID,
			&bannerContent.Content,
			&bannerContent.Version,
			&bannerContent.CreatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "can't scan get contents for banner query result")
		}

		banners[bannerIndexes[bannerID]].Versions = append(banners[bannerIndexes[bannerID]].Versions, bannerContent)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "can't end scan get contents for banner query result")
	}

	return banners, nil
}

func (st *storageData) GetBanners(featureID, tagID uint32, offset, limit uint64) ([]models.Banner, error) {
	var banners []models.Banner
	temp := models.BannerInfo{
		FeatureID: featureID,
		TagID:     tagID,
	}

	if err := WithTransaction(st.db,
		func(tx *sql.Tx) error {
			var err error

			banners, err = filterBanners(tx, &temp, offset, limit)
			if err != nil {
				return err
			}

			if len(banners) == 0 {
				return nil
			}

			banners, err = selectTagFeatureForBanners(tx, banners)
			if err != nil {
				return err
			}

			banners, err = selectContentForBanners(tx, banners)
			if err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, errors.Wrapf(err,
			"when selecting banners for admin with feature id %d, tag id %d, limit %d and offset %d",
			temp.FeatureID, temp.TagID, limit, offset)
	}

	return banners, nil
}

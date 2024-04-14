package storage

const (
	getQuery = `
		SELECT vb.content FROM banner
		   INNER JOIN features_tags_banner on (features_tags_banner.banner_id = banner.id and not deleted)
		   LEFT JOIN version_banner as vb on (vb.banner_id = banner.id)
		WHERE is_active and vb.version = COALESCE($3::bigint, banner.last_version) 
		  		and feature_id = $1 and tag_id = $2 LIMIT 1
	`

	createQuery = `
		INSERT INTO banner (is_active)
		VALUES ($1)
		RETURNING id
	`

	addFeaturesAndTagsQuery = `
		INSERT INTO features_tags_banner (banner_id, feature_id, tag_id)
		SELECT $1, $2, tag
		FROM unnest($3::bigint[]) as tag		
	`

	addContentQuery = `
		INSERT INTO version_banner (banner_id, content) VALUES ($1, $2)
	`

	getTagQuery = `
		SELECT banner_id, array_agg(tag_id), feature_id FROM features_tags_banner 
			WHERE banner_id = ANY ($1::bigint[])
		GROUP BY banner_id, feature_id 
	`

	getVersionQuery = `
		SELECT banner_id, content, version, created_at FROM version_banner WHERE banner_id = ANY ($1::bigint[])
	`

	filterNotNullQuery = `
		SELECT DISTINCT banner.id, is_active,
						created_at, updated_at FROM banner
			INNER JOIN features_tags_banner as ftb ON (ftb.banner_id = banner.id  and not deleted)
			WHERE (CASE WHEN $1::bigint IS NOT NULL THEN feature_id = $1 ELSE true END)
			and (CASE WHEN $2::bigint IS NOT NULL THEN tag_id = $2 ELSE true END)
			LIMIT $3 OFFSET $4
	`

	deleteQuery = `
		DELETE FROM banner WHERE id IN 
		(SELECT DISTINCT banner_id FROM features_tags_banner 
				WHERE not deleted and banner_id = $1) RETURNING id
	`

	checkDeleted = `
		SELECT banner_id FROM features_tags_banner WHERE banner_id = $1 and not deleted
	`

	updateActiveQuery = `
		UPDATE banner SET is_active = $2
			WHERE id = $1
			RETURNING id
	`

	updateFeaturesQuery = `
		UPDATE features_tags_banner SET feature_id = $2 WHERE banner_id = $1
	`

	deleteFeaturesTagsQuery = `
		WITH deleted_features AS (
			DELETE FROM features_tags_banner WHERE banner_id = $1 RETURNING feature_id
		)
		SELECT DISTINCT feature_id FROM deleted_features LIMIT 1
	`

	filterNullQuery = `
		SELECT banner.id, is_active, created_at, updated_at FROM banner LIMIT $1 OFFSET $2
	`
)

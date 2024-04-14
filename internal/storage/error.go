package storage

import "github.com/pkg/errors"

const (
	uniqueConflictCode   = "23505"
	uniqueConstraintName = "banner_identifier"
)

var (
	ErrorBannerNotFound       = errors.New("banner not found")
	ErrorBannerConflictExists = errors.New("banner with presented pair featured id and tag id is already exists")
)

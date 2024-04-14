package models

import (
	"encoding/json"
	"time"
)

type BannerUpdate struct {
	Content   json.RawMessage
	FeatureID uint32
	TagIDs    []uint32
	IsActive  bool
}
type BannerInfo struct {
	FeatureID uint32
	TagID     uint32
}

type Content struct {
	Version   uint32
	Content   json.RawMessage
	CreatedAt time.Time
}

type Banner struct {
	ID        uint32
	FeatureID uint32
	TagIDs    []uint32
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Versions  []Content
}

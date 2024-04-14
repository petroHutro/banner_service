package storage

import (
	"banner_service/internal/models"
	"database/sql"
)

type User interface {
	GetBanner(featureID, tagID uint32) (string, error)
}

type Admin interface {
	GetBanners(featureID, tagID uint32, offset, limit uint64) ([]models.Banner, error)
	CreateBanner(featureID uint32, tagIDs []uint32, content string, isActive bool) (uint32, error)
	DeleteBanner(id uint32) error
	UpdateBanner(id uint32, banner *models.BannerUpdate) error
}

type Storage interface {
	User
	Admin
}

type storageData struct {
	db *sql.DB
}

package handlers

import (
	"banner_service/internal/authorization"
	"banner_service/internal/caches"
	"banner_service/internal/storage"
	"encoding/json"
)

type Handler struct {
	storage             storage.Storage
	cache               *caches.Cache
	authorization       *authorization.Authorization
	errorIncorrectData  []byte
	errorInternalServer []byte
}

type ErrorBannerModel struct {
	Error string `json:error`
}

func Init(storage storage.Storage, cache *caches.Cache, authorization *authorization.Authorization) *Handler {
	errorIncorrectData, _ := json.Marshal(ErrorBannerModel{Error: "incorrect data"})

	errorInternalServer, _ := json.Marshal(ErrorBannerModel{Error: "internal server error"})

	return &Handler{
		storage:             storage,
		cache:               cache,
		authorization:       authorization,
		errorIncorrectData:  errorIncorrectData,
		errorInternalServer: errorInternalServer,
	}
}

package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

type BannerContentModel struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

type GetBannerModel struct {
	TagId     uint32 `schema:"tag_id,omitempty"`
	FeatureId uint32 `schema:"feature_id,omitempty"`
	Limit     uint64 `schema:"limit,omitempty"`
	Offset    uint64 `schema:"offset,omitempty"`
}

type ResponseBannerModel struct {
	BannerID  uint32             `json:"banner_id"`
	TagIDs    []uint32           `json:"tag_ids"`
	FeatureID uint32             `json:"feature_id"`
	IsActive  bool               `json:"is_active"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Content   BannerContentModel `json:"content"`
}

func (h *Handler) GetBanner(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	bannerModel := new(GetBannerModel)
	if err := schema.NewDecoder().Decode(bannerModel, r.Form); err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	// TO DO
	respBannerModel, err := h.storage.GetBanners(bannerModel.FeatureId, bannerModel.TagId, 1, 10)
	if err != nil {
		logger.Error("error working with the database: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(respBannerModel); err != nil {
		logger.Error("error forming the response body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger.Info("handler /banner worked correctly, the status 200 was received")

	w.WriteHeader(http.StatusOK)
}

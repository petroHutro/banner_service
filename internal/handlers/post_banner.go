package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"
	"net/http"
)

type PostBannerModel struct {
	TagIds    []uint32           `json:"tag_ids,omitempty"`
	FeatureId uint32             `json:"feature_id,omitempty"`
	IsActive  bool               `json:"is_active,omitempty"`
	Content   BannerContentModel `json:"content,omitempty"`
}

type ResponsePostBannerModel struct {
	BannerID uint32 `json:"banner_id"`
}

func (h *Handler) PostBanner(w http.ResponseWriter, r *http.Request) {
	var addBannerModel PostBannerModel

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&addBannerModel); err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	jsonData, err := json.Marshal(addBannerModel.Content)
	if err != nil {
		logger.Error("content request data format: %v", err)
		return
	}

	respBannerModel, err := h.storage.CreateBanner(addBannerModel.FeatureId, addBannerModel.TagIds, string(jsonData), addBannerModel.IsActive)
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
	w.WriteHeader(http.StatusCreated)
}

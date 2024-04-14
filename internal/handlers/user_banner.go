package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"
	"errors"

	"net/http"

	"github.com/gorilla/schema"
)

var ErrBannerNotFound = errors.New("the banner with the specified data was not found")

type GetUserBannerModel struct {
	TagId           uint32 `schema:"tag_id"`
	FeatureId       uint32 `schema:"feature_id"`
	UseLastRevision bool   `schema:"use_last_revision,omitempty"`
}

func (h *Handler) UserBanner(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	userBannerModel := new(GetUserBannerModel)
	if err := schema.NewDecoder().Decode(userBannerModel, r.Form); err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	bannerContentModel, err := h.storage.GetBanner(userBannerModel.FeatureId, userBannerModel.TagId)
	if err != nil && !errors.Is(err, ErrBannerNotFound) {
		logger.Error("error working with the database: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	} else if errors.Is(err, ErrBannerNotFound) {
		logger.Error("no data was found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(bannerContentModel); err != nil {
		logger.Error("error forming the response body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

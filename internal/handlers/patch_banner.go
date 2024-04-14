package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type PatchBannerModel struct {
	TagIds    []uint32           `json:"tag_ids,omitempty"`
	FeatureId uint32             `json:"feature_id,omitempty"`
	Content   BannerContentModel `json:"content,omitempty"`
	IsActive  bool               `json:"is_active,omitempty"`
	Id        uint32
}

func (h *Handler) PatchBanner(w http.ResponseWriter, r *http.Request) {
	var patchBannerModel PatchBannerModel

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&patchBannerModel); err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	stringBannerId := strings.Split(r.URL.Path, "/")[2]
	bannerId, err := strconv.Atoi(stringBannerId)
	if err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}
	patchBannerModel.Id = uint32(bannerId)

	err = h.storage.UpdateBanner(patchBannerModel.Id, nil)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

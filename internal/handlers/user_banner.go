package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"

	"net/http"
)

type DataLogin struct {
	Tag      int  `json:"tag_id"`
	Feature  int  `json:"feature_id"`
	Revision bool `json:"use_last_revision"`
}

func (h *Handler) UserBanner(w http.ResponseWriter, r *http.Request) {
	var data DataLogin

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logger.Error("bad json: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"

	"net/http"
)

type DataLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) UserBanner(w http.ResponseWriter, r *http.Request) {
	var data DataLogin

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		// logger.Error("bad json: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

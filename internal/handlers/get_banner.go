package handlers

import (
	"net/http"
)

func (h *Handler) GetBanner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

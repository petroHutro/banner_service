package handlers

import (
	"net/http"
)

func (h *Handler) PostBanner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

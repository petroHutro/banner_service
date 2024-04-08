package handlers

import (
	"net/http"
)

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"net/http"
)

func (h *Handler) PatchBanner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

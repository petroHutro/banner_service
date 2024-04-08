package handlers

import (
	"net/http"
)

func (h *Handler) Banner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

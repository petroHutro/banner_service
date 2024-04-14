package handlers

import (
	"banner_service/internal/logger"
	"encoding/json"
	"net/http"
	"time"
)

type rToken struct {
	Token string
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TO DO getToken(accountName)
	accountName := r.URL.Query().Get("accountName")
	token, err := h.authorization.BuildJWTString(accountName, time.Hour*3)
	if err != nil {
		logger.Error("error when creating a token during authorization for user %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(rToken{Token: token})
	if err != nil {
		logger.Error("cannot json to byte: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

package handlers

import (
	"banner_service/internal/logger"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	stringIdBanner := strings.Split(r.URL.Path, "/")[2]
	idBanner, err := strconv.Atoi(stringIdBanner)
	if err != nil {
		logger.Error("invalid request data format: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	err = h.storage.DeleteBanner(uint32(idBanner))
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
	w.WriteHeader(http.StatusNoContent)
	logger.Info("handler /banner/{id} worked correctly, the status 200 was received")
}

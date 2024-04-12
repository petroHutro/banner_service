package app

import (
	"banner_service/internal/config"
	"banner_service/internal/handlers"
	"banner_service/internal/logger"
	"banner_service/internal/router"
	"banner_service/internal/storage"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type App struct {
	conf    *config.Flags
	router  *chi.Mux
	handler *handlers.Handler
	storage storage.Storage
}

func newApp() (*App, error) {
	conf := config.NewFlags()

	router := router.CreateRouter()

	handler := handlers.Init()

	storage, err := storage.InitStorage(&conf.Storage)
	if err != nil {
		return nil, fmt.Errorf("cannot init storage: %w", err)
	}

	logger.Info("Running server: address:%s port:%d", conf.Host, conf.Port)

	return &App{
		conf:    &conf,
		router:  router,
		handler: handler,
		storage: storage,
	}, nil
}

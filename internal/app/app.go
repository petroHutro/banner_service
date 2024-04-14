package app

import (
	"banner_service/internal/authorization"
	"banner_service/internal/caches"
	"banner_service/internal/config"
	"banner_service/internal/handlers"
	"banner_service/internal/logger"
	"banner_service/internal/router"
	"banner_service/internal/storage"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type App struct {
	conf          *config.Flags
	router        *chi.Mux
	handler       *handlers.Handler
	storage       storage.Storage
	cache         *caches.Cache
	authorization *authorization.Authorization
}

func newApp() (*App, error) {
	conf := config.NewFlags()

	if err := logger.Init(conf.Logger); err != nil {
		return nil, fmt.Errorf("cannot init logger: %w", err)
	}

	router := router.Init()

	storage, err := storage.Init(&conf.Storage)
	if err != nil {
		return nil, fmt.Errorf("cannot init storage: %w", err)
	}

	cache, err := caches.Init(conf.Url)
	if err != nil {
		return nil, fmt.Errorf("cannot init caches: %w", err)
	}

	authorization := authorization.Init(conf.TokenSecretKey)

	handler := handlers.Init(storage, cache, authorization)

	logger.Info("Running server: address:%s port:%d", conf.Host, conf.Port)

	return &App{
		conf:          &conf,
		router:        router,
		handler:       handler,
		storage:       storage,
		cache:         cache,
		authorization: authorization,
	}, nil
}

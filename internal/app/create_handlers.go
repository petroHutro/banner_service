package app

import (
	"banner_service/internal/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) createMiddlewareHandlers() {
	a.router.Use(logger.LoggingMiddleware)
}

func (a *App) createHandlers() {
	a.router.Route("/user_banner", func(r chi.Router) {
		r.With(a.authorization.AuthorizationMiddleware(),
			a.cache.CacheMiddleware).
			Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.handler.UserBanner(w, r)
			})
	})

	a.router.Route("/banner", func(r chi.Router) {
		r.With(a.authorization.AuthorizationMiddleware()).
			Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.handler.GetBanner(w, r)
			})

		r.With(a.authorization.AuthorizationMiddleware()).
			Post("/", func(w http.ResponseWriter, r *http.Request) {
				a.handler.PostBanner(w, r)
			})

		r.With(a.authorization.AuthorizationMiddleware()).
			Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
				a.handler.PatchBanner(w, r)
			})

		r.With(a.authorization.AuthorizationMiddleware()).
			Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
				a.handler.DeleteBanner(w, r)
			})
	})

	a.router.Route("/login", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			a.handler.Login(w, r)
		})
	})
}

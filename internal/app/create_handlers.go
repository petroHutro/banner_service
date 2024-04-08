package app

import (
	"banner_service/internal/authorization"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) createMiddlewareHandlers() {
	// a.router.Use(logger.LoggingMiddleware)
}

func (a *App) createHandlers() {
	a.router.Route("/user_banner", func(r chi.Router) {
		r.With(authorization.AuthorizationMiddleware("")).
			Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.handler.UserBanner(w, r)
			})
	})

	a.router.Route("/banner", func(r chi.Router) {
		r.With(authorization.AuthorizationMiddleware("")).
			Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.handler.Banner(w, r)
			})

		r.With(authorization.AuthorizationMiddleware("")).
			Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
				a.handler.PatchBanner(w, r)
			})

		r.With(authorization.AuthorizationMiddleware("")).
			Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
				a.handler.DeleteBanner(w, r)
			})
	})
}

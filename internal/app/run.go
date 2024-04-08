package app

import (
	"fmt"
	"net/http"
	"strconv"
)

func Run() error {
	app, err := newApp()
	if err != nil {
		return fmt.Errorf("cannot init app: %w", err)
	}

	app.createMiddlewareHandlers()
	app.createHandlers()

	address := app.conf.Host + ":" + strconv.Itoa(app.conf.Port)

	return http.ListenAndServe(address, app.router)
}

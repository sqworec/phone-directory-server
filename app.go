package main

import (
	"phone-directory-server/data"
	"phone-directory-server/handlers"
)

type App struct {
	h *handlers.HTTPHandler
}

func NewApp(c Config) *App {
	db := data.NewDB(c.DSN)
	h := handlers.NewHTTPHandler(db)

	return &App{
		h: h,
	}
}

func (app *App) Start(port int) {
	app.h.ServeHTTP(port)
}

package handler

import (
	"github.com/maruki00/deligo/internal/profile/app"
)

type Handler struct {
	app *app.App
}

func New(instance *app.App) *Handler {
	return &Handler{app: instance}
}

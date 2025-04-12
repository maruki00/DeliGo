package controllers

import (
	"delivery/internal/auth/app"
)

type CustomerController struct {
	app *app.App
}

func NewAuthController(app *app.App) *CustomerController {
	return &CustomerController{
		app: app,
	}
}

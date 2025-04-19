package controllers

import (
	"deligo/internal/auth/app"
)

type CustomerController struct {
	app *app.App
}

func NewAuthController(app *app.App) *CustomerController {
	return &CustomerController{
		app: app,
	}
}

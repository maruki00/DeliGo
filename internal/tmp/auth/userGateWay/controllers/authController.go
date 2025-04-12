package controllers

import (
	"delivery/internal/auth/app"
	"delivery/internal/auth/domain/dtos"
	"delivery/internal/auth/userGateWay/requests"
	"delivery/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	app *app.App
}

func NewAuthController(app *app.App) *AuthController {
	return &AuthController{
		app: app,
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &requests.LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.app.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	result := obj.app.InputPort.Login(dtos.LoginDTO{
		Login:    request.Login,
		Password: request.Password,
	})

	ctx.JSON(result.GetResponse().Status, result.GetResponse())
}

func (obj AuthController) Register(ctx *gin.Context) {

	request := &requests.RegisterRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.app.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	res := obj.app.InputPort.Register(dtos.RegisterDTO{
		FullName:  request.FullName,
		UserName:  request.UserName,
		Email:     request.Email,
		Address:   request.Address,
		Password:  request.Password,
		UserLevel: 0,
		UserType:  "costumer",
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj AuthController) TwoFactoryConfirm(ctx *gin.Context) {
	request := &requests.TwoFactoryConfirmRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.app.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	res := obj.app.InputPort.TwoFactoryConfirm(dtos.TwoFactoryConfirmDTO{
		Email: request.Email,
		Pin:   request.Pin,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())

}

func (obj *AuthController) Logout(ctx *gin.Context) {
	request := &requests.LogoutRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}

	if err := obj.app.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}
	obj.app.InputPort.Logout(dtos.LogoutDTO{
		Token: request.Token,
	})
	utils.Success(ctx, http.StatusNoContent, "Success", nil)
}

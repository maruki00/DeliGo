package shared_core

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(ctx *gin.Context, Validate *validator.Validate, request interface{}) error {
	if err := ctx.BindJSON(request); err != nil {
		return err
	}

	if err := Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())

		return fmt.Errorf(errorMessage)
	}
	return nil
}

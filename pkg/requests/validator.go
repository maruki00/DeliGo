package pkgValidator

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"

// 	pkgValidations "github.com/maruki00/deligo/pkg/validations"

// 	"github.com/go-playground/validator/v10"
// )

// type Request struct {
// 	v *validator.Validate
// }

// func NewRequest() *Request {
// 	val := validator.New()
// 	val.RegisterValidation("boolean", pkgValidations.BooleanValidator)
// 	return &Request{
// 		v: val,
// 	}
// }

// func (r *Request) BindJson(data []byte, out interface{}) error {
// 	if err := json.Unmarshal(data, out); err != nil {
// 		return errors.New("could not unmarshal the request body")
// 	}
// 	return nil
// }

// func (r *Request) Validated(req interface{}) error {
// 	if err := r.v.Struct(req); err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
// 		return errors.New(errorMessage)
// 	}
// 	return nil
// }

// func Validate(ctx context.Context, Validate *validator.Validate, request interface{}) error {

// 	if err := Validate.Struct(request); err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())

// 		return fmt.Errorf(errorMessage)
// 	}
// 	return nil
// }

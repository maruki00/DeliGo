package command

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type UpdateProductCommand struct {
	ID sharedvo.ID
}

func (_this *UpdateProductCommand) Name() string {
	return "UpdateProductCommand"
}

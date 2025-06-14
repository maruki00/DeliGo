package command

import sharedvo "deligo/internal/shared/valueobject"

type UpdateProductCommand struct {
	ID sharedvo.ID
}

func (_this *UpdateProductCommand) Name() string {
	return "UpdateProductCommand"
}

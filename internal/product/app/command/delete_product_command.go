package command

import sharedvo "deligo/internal/shared/valueobject"

type DeleteProductCommand struct {
	ID sharedvo.ID
}

func (_this *DeleteProductCommand) Name() string {
	return "DeleteProductCommand"
}

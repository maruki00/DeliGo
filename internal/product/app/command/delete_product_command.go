package command

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type DeleteProductCommand struct {
	ID sharedvo.ID
}

func (_this *DeleteProductCommand) Name() string {
	return "DeleteProductCommand"
}

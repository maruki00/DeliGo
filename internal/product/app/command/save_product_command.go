package command

import sharedvo "github.com/maruki00/deligo/internal/shared/valueobject"

type SaveProductCommand struct {
	ID    sharedvo.ID
	Label string
	Price float32
	Qty   uint32
}

func (_this *SaveProductCommand) Name() string {
	return "SaveProductCommand"
}

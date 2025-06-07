package command

import (
	sharedvo "deligo/internal/shared/valueobject"
)

type AcceptOrderCommand struct {
	ID sharedvo.ID
}

func (_this *AcceptOrderCommand) Name() string {
	return "AcceptOrderCommand"
}

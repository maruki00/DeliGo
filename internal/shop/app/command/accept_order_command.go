package command

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type AcceptOrderCommand struct {
	ID sharedvo.ID
}

func (_this *AcceptOrderCommand) Name() string {
	return "AcceptOrderCommand"
}

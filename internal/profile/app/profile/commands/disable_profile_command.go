package commands

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
)

type DiscableProfileCommand struct {
	ID valueobjects.ID
}

func (_this *DiscableProfileCommand) Name() string {
	return "DiscableProfileCommand"
}

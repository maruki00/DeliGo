package command

import (
	shared_valueobject "github.com/maruki00/deligo/internal/shared/value_object"
)

type DiscableProfileCommand struct {
	ID shared_valueobject.ID
}

func (_this *DiscableProfileCommand) Name() string {
	return "DiscableProfileCommand"
}

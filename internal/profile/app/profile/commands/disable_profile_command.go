package commands

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
)

type DiscableProfileCommand struct {
	ID shared_valueobject.ID
}

func (_this *DiscableProfileCommand) Name() string {
	return "DiscableProfileCommand"
}

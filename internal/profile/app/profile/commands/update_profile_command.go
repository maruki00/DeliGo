package commands

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
)

type UpdateProfileCommand struct {
	ID     shared_valueobject.ID
	Fields map[string]any
}

func (_this *UpdateProfileCommand) Name() string {
	return "UpdateProfileCommand"
}

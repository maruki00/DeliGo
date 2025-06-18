package command

import (
	shared_valueobject "github.com/maruki00/deligo/internal/shared/value_object"
)

type UpdateProfileCommand struct {
	ID     shared_valueobject.ID
	Fields map[string]any
}

func (_this *UpdateProfileCommand) Name() string {
	return "UpdateProfileCommand"
}

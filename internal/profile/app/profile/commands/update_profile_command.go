package commands

import "github.com/google/uuid"

type UpdateProfileCommand struct {
	ID     uuid.UUID
	Fields map[string]any
}

func (_this *UpdateProfileCommand) Name() string {
	return "UpdateProfileCommand"
}

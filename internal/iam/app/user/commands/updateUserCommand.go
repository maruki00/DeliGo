package userCommands

import "github.com/google/uuid"

type UpdateUserCommand struct {
	ID     uuid.UUID         `json:"id"`
	Fields map[string]string `json:"fields"`
}

func (_this *UpdateUserCommand) CommandName() string {
	return "UpdateUserCommand"
}

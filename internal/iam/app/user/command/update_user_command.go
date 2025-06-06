package userCommand

import "github.com/google/uuid"

type UpdateUserCommand struct {
	ID     uuid.UUID              `json:"id"`
	Fields map[string]interface{} `json:"fields"`
}

func (_this *UpdateUserCommand) Name() string {
	return "UpdateUserCommand"
}

package userCommands

import "github.com/google/uuid"

type DeleteUserCommand struct {
	ID uuid.UUID
}

func (_this *DeleteUserCommand) CommandName() string {
	return "DeleteUserCommand"
}

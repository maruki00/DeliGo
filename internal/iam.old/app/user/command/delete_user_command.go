package userCommand

import "github.com/google/uuid"

type DeleteUserCommand struct {
	ID uuid.UUID
}

func (_this *DeleteUserCommand) Name() string {
	return "DeleteUserCommand"
}

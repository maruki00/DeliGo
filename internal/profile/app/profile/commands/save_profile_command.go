package commands

import "github.com/google/uuid"

type SaveProfileCommand struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	FullName string
	Avatar   string
	Bio      string
}

func (_this *SaveProfileCommand) Name() string {
	return "SaveProfileCommand"
}

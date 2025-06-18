package command

import (
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type SaveProfileCommand struct {
	ID       sharedvo.ID
	UserID   sharedvo.ID
	FullName string
	Avatar   string
	Bio      string
}

func (_this *SaveProfileCommand) Name() string {
	return "SaveProfileCommand"
}

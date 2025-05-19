package commands

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
)

type SaveProfileCommand struct {
	ID       shared_valueobject.ID
	UserID   shared_valueobject.ID
	FullName string
	Avatar   string
	Bio      string
}

func (_this *SaveProfileCommand) Name() string {
	return "SaveProfileCommand"
}

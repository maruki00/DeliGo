package commands

import shared_valueobject "github.com/maruki00/deligo/internal/shared/domain/sharedvo"

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

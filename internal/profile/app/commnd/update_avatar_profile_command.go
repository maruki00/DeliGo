package commands

import (
	shared_valueobject "github.com/maruki00/deligo/internal/shared/valueobject"
)

type UpdateProfileAvatarCommand struct {
	ID     shared_valueobject.ID
	Avatar string
}

func (_this *UpdateProfileAvatarCommand) Name() string {
	return "UpdateProfileAvatarCommand"
}

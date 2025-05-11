package commands

import "github.com/google/uuid"

type UpdateProfileAvatarCommand struct {
	ID uuid.UUID
}

func (_this *UpdateProfileAvatarCommand) Name() string {
	return "UpdateProfileAvatarCommand"
}

package commands

import "github.com/google/uuid"

type DiscableProfileCommand struct {
	ID uuid.UUID
}

func (_this *DiscableProfileCommand) Name() string {
	return "DiscableProfileCommand"
}

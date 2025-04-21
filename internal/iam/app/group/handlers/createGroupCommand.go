package groupHandlers

import (
	"time"

	"github.com/google/uuid"
)

type CreateGroupCommand struct {
	ID                uuid.UUID
	Username          string
	Email             string
	Password          string
	PasswordChangedAt *time.Time
	IsActive          bool
	MFAEnabled        bool
	MFASecret         string
}

func (_this *CreateGroupCommand) CommandName() string {
	return "CreateGroupCommand"
}

package userCommand

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserCommand struct {
	ID                uuid.UUID
	Username          string
	Email             string
	Password          string
	PasswordChangedAt *time.Time
	IsActive          bool
	MFAEnabled        bool
	MFASecret         string
}

func (_this *CreateUserCommand) Name() string {
	return "CreateUserCommand"
}

package userHandlers

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	"deligo/internal/iam/domain/contracts"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	pkgCqrs "deligo/pkg/cqrs"
)

type CreateUserHandler struct {
	userRepo contracts.IUserRepository
}

func NewCreateUserHandler(userRepo contracts.IUserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		userRepo: userRepo,
	}
}

func (_this *CreateUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommands.CreateUserCommand)
	err := _this.userRepo.Save(ctx, &models.User{
		ID:                valueobjects.ID(cmd.ID.String()),
		Username:          cmd.Username,
		Email:             cmd.Email,
		Password:          valueobjects.Password(cmd.Password),
		PasswordChangedAt: cmd.PasswordChangedAt,
		IsActive:          cmd.IsActive,
		MFAEnabled:        cmd.MFAEnabled,
		MFASecret:         cmd.MFASecret,
	})
	return err
}

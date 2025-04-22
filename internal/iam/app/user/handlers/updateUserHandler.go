package userHandlers

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type UpdateUserHandler struct {
	userRepo contracts.IUserRepository
}

func NewUpdateUserHandler(userRepo contracts.IUserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{
		userRepo: userRepo,
	}
}

func (_this *UpdateUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommands.UpdateUserCommand)
	// TODO : Implement the logic to update the user
	err := _this.userRepo.Update(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

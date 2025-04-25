package userHandlers

import (
	"context"
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
	// cmd := command.(*userCommands.UpdateUserCommand)
	// id := cmd.ID
	// err := _this.userRepo.Update(ctx, id, cmd.Fields)
	// if err != nil {
	// 	return err
	// }
	return nil
}

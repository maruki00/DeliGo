package userHandlers

import (
	"context"
	"fmt"

	userCommands "deligo/internal/iam/app/user/commands"
	"deligo/internal/iam/domain/contracts"
	valueobjects "deligo/internal/iam/domain/valueobject"
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
	fmt.Println("called : func (_this *UpdateUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error ")
	cmd := command.(*userCommands.UpdateUserCommand)
	err := _this.userRepo.Update(ctx, valueobjects.ID(cmd.ID.String()), cmd.Fields)
	if err != nil {
		return err
	}
	return nil
}

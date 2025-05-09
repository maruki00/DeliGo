package userHandlers

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	"deligo/internal/iam/domain/contracts"
	valueobjects "deligo/internal/iam/domain/valueobject"
	pkgCqrs "deligo/pkg/cqrs"
	"fmt"
)

type DeleteUserHandler struct {
	userRepo contracts.IUserRepository
}

func NewDeleteUserHandler(userRepo contracts.IUserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		userRepo: userRepo,
	}
}

func (_this *DeleteUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommands.DeleteUserCommand)
	fmt.Println("hello worldf")
	err := _this.userRepo.Delete(ctx, valueobjects.ID(cmd.ID.String()))
	return err
}

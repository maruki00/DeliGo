package userHandler

import (
	"context"
	userCommand "github.com/maruki00/deligo/internal/iam/app/user/command"
	"github.com/maruki00/deligo/internal/iam/domain/contracts"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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
	cmd := command.(*userCommand.DeleteUserCommand)
	fmt.Println("hello worldf")
	err := _this.userRepo.Delete(ctx, valueobjects.ID(cmd.ID.String()))
	return err
}

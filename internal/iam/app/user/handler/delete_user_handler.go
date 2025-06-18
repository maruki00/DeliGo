package userHandler

import (
	"context"
	"fmt"

	userCommand "github.com/maruki00/deligo/internal/iam/app/user/command"
	"github.com/maruki00/deligo/internal/iam/domain/contract"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type DeleteUserHandler struct {
	userRepo contract.IUserRepository
}

func NewDeleteUserHandler(userRepo contract.IUserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		userRepo: userRepo,
	}
}

func (_this *DeleteUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommand.DeleteUserCommand)
	fmt.Println("hello worldf")
	err := _this.userRepo.Delete(ctx, sharedvo.ID(cmd.ID.String()))
	return err
}

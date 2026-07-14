package userHandler

import (
	"context"
	"fmt"

	userCommand "github.com/maruki00/deligo/internal/iam/app/user/command"
	"github.com/maruki00/deligo/internal/iam/domain/contract"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateUserHandler struct {
	userRepo contract.IUserRepository
}

func NewUpdateUserHandler(userRepo contract.IUserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{
		userRepo: userRepo,
	}
}

func (_this *UpdateUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommand.UpdateUserCommand)
	err := _this.userRepo.Update(ctx, sharedvo.ID(cmd.ID.String()), cmd.Fields)
	if err != nil {
		return fmt.Errorf("error : %s", err.Error())
	}
	return nil
}

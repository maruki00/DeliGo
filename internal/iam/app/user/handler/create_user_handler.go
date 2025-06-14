package userHandler

import (
	"context"
	userCommand "github.com/maruki00/deligo/internal/iam/app/user/command"
	"github.com/maruki00/deligo/internal/iam/domain/contracts"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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
	cmd := command.(*userCommand.CreateUserCommand)
	pwd, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		return err
	}

	err = _this.userRepo.Save(ctx, &model.User{
		ID:                valueobjects.ID(cmd.ID.String()),
		Username:          cmd.Username,
		Email:             cmd.Email,
		Password:          pwd,
		PasswordChangedAt: nil,
		IsActive:          false,
		MFAEnabled:        false,
		MFASecret:         "",
	})
	return err
}

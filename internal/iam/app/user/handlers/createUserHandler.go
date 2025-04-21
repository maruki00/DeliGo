package userHandlers

import (
	"context"
	"deligo/internal/iam/domain/contracts"
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
	return nil
}

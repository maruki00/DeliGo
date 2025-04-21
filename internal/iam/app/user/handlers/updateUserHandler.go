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
	return nil
}

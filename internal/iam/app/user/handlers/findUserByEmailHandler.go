package userHandlers

import (
	"context"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type FindUserByEmailHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByEmailHandler(userRepo contracts.IUserRepository) *FindUserByEmailHandler {
	return &FindUserByEmailHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByEmailHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	return nil
}

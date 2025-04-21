package userHandlers

import (
	"context"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type FindUserByUsernameHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByUsernameHandler(userRepo contracts.IUserRepository) *FindUserByUsernameHandler {
	return &FindUserByUsernameHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByUsernameHandler) Handle(ctx context.Context, command pkgCqrs.Query) (interface{}, error) {
	return nil, nil
}

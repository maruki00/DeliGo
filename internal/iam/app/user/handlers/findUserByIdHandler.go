package userHandlers

import (
	"context"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type FindUserByIdHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByIdHandler(userRepo contracts.IUserRepository) *FindUserByIdHandler {
	return &FindUserByIdHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByIdHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	return nil, nil
}

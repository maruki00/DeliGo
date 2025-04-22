package userHandlers

import (
	"context"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type ListUsersByTenantHandler struct {
	userRepo contracts.IUserRepository
}

func NewListUsersByTenantHandler(userRepo contracts.IUserRepository) *ListUsersByTenantHandler {
	return &ListUsersByTenantHandler{
		userRepo: userRepo,
	}
}

func (_this *ListUsersByTenantHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	return nil, nil
}

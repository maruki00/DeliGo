package handlers

import (
	"context"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type CreateShopHHandler struct {
	repo contracts.IShopRepository
}

func NewCreateShopHandler(repo contracts.IShopRepository) *CreateShopHHandler {
	return &CreateShopHHandler{
		repo: repo,
	}
}

func (_this *CreateShopHHandler) handler(ctx context.Context, command *pkgCqrs.Command) error {

	return nil
}

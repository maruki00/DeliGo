package handlers

import (
	"context"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type UpdateShopHandler struct {
	repo contracts.IShopRepository
}

func NewUpdateShopHandler(repo contracts.IShopRepository) *UpdateShopHandler {
	return &UpdateShopHandler{
		repo: repo,
	}
}

func (_this *UpdateShopHandler) handler(ctx context.Context, command *pkgCqrs.Command) error {

	return nil
}

package handlers

import (
	"context"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type CloseShopHandler struct {
	repo contracts.IShopRepository
}

func NewCloseShopHandler(repo contracts.IShopRepository) *CloseShopHandler {
	return &CloseShopHandler{
		repo: repo,
	}
}

func (_this *CloseShopHandler) handler(ctx context.Context, command *pkgCqrs.Command) error {

	return nil
}

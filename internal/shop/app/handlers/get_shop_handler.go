package handlers

import (
	"context"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetShopHandler struct {
	repo contracts.IShopRepository
}

func NewGetShopHandler(repo contracts.IShopRepository) *GetShopHandler {
	return &GetShopHandler{
		repo: repo,
	}
}

func (_this *GetShopHandler) handler(ctx context.Context, query *pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}

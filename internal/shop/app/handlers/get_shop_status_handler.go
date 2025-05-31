package handlers

import (
	"context"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetShopStatusHandler struct {
	repo contracts.IShopRepository
}

func NewGetShopStatusHandler(repo contracts.IShopRepository) *GetShopStatusHandler {
	return &GetShopStatusHandler{
		repo: repo,
	}
}

func (_this *GetShopStatusHandler) handler(ctx context.Context, query *pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}

package handler

import (
	"context"
	"deligo/internal/shop/app/query"
	"deligo/internal/shop/domain/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetShopHandler struct {
	repo contract.IShopRepository
}

func NewGetShopHandler(repo contract.IShopRepository) *GetShopHandler {
	return &GetShopHandler{
		repo: repo,
	}
}

func (_this *GetShopHandler) handler(ctx context.Context, qry pkgCqrs.Query) (interface{}, error) {
	q := qry.(*query.GetShopSquery)
	return _this.repo.GetByID(ctx, q.ID)
}

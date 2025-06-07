package handler

import (
	"context"
	"deligo/internal/shop/app/query"
	"deligo/internal/shop/domain/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetShopStatusHandler struct {
	repo contract.IShopRepository
}

func NewGetShopStatusHandler(repo contract.IShopRepository) *GetShopStatusHandler {
	return &GetShopStatusHandler{
		repo: repo,
	}
}

func (_this *GetShopStatusHandler) handler(ctx context.Context, qry pkgCqrs.Query) (interface{}, error) {
	q := qry.(*query.GetShopStatusQuery)
	return _this.repo.GetShopStatus(ctx, q.ID)
}

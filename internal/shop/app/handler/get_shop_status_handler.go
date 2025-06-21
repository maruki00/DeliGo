package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/shop/app/query"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type GetShopStatusHandler struct {
	repo contract.IShopRepository
}

func NewGetShopStatusHandler(repo contract.IShopRepository) *GetShopStatusHandler {
	return &GetShopStatusHandler{
		repo: repo,
	}
}

func (_this *GetShopStatusHandler) handle(ctx context.Context, qry pkgCqrs.Query) (interface{}, error) {
	q := qry.(*query.GetShopStatusQuery)
	return _this.repo.GetShopStatus(ctx, q.ID)
}

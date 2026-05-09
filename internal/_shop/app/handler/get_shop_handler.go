package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/shop/app/query"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type GetShopHandler struct {
	repo contract.IShopRepository
}

func NewGetShopHandler(repo contract.IShopRepository) *GetShopHandler {
	return &GetShopHandler{
		repo: repo,
	}
}

func (_this *GetShopHandler) handle(ctx context.Context, qry pkgCqrs.Query) (interface{}, error) {
	q := qry.(*query.GetShopSquery)
	return _this.repo.GetByID(ctx, q.ID)
}

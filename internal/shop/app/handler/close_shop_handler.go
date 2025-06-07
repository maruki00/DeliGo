package handler

import (
	"context"
	"deligo/internal/shop/app/queries"
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

func (_this *CloseShopHandler) handler(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {

	qry := query.(*queries.GetShopSquery)
	return _this.repo.GetByID(ctx, qry.ID)
}

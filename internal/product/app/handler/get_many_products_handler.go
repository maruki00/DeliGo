package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/product/app/query"
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type GetManyProductsHandler struct {
	repo contract.IProductRepository
}

func NewGetManyProductsHandler(repo contract.IProductRepository) GetManyProductsHandler {
	return GetManyProductsHandler{
		repo: repo,
	}
}

func (_this *GetManyProductsHandler) handle(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	qry := q.(*query.GetManyProductQuery)

	res, err := _this.repo.Get
	return nil, nil
}

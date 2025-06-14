package handler

import (
	"context"
	"deligo/internal/product/domian/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetManyProductsHandler struct {
	repo contract.IProductRepository
}

func NewGetManyProductsHandler(repo contract.IProductRepository) GetManyProductsHandler {
	return GetManyProductsHandler{
		repo: repo,
	}
}

func (_this *GetManyProductsHandler) handler(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}

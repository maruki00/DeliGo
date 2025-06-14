package handler

import (
	"context"
	"deligo/internal/product/domian/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetProductByMultipleIDHandler struct {
	repo contract.IProductRepository
}

func NewGetProductByMultipleIDHandler(repo contract.IProductRepository) GetProductByMultipleIDHandler {
	return GetProductByMultipleIDHandler{
		repo: repo,
	}
}

func (_this *GetProductByMultipleIDHandler) handler(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}

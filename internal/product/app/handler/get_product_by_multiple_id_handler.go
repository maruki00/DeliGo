package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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

package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type GetProductBy
IDHandler struct {
	repo contract.IProductRepository
}

func NewGetProductByIDHandler(repo contract.IProductRepository) GetProductByIDHandler {
	return GetProductBy
IDHandler{
		repo: repo,
	}
}

func (_this *GetProductByIDHandler) handler(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}

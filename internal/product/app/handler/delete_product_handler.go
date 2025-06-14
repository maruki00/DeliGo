package handler

import (
	"deligo/internal/product/domian/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type DeleteProductHandler struct {
	repo contract.IProductRepository
}

func NewDeleteProductHandler(repo contract.IProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{
		repo: repo,
	}
}

func (_this *DeleteProductHandler) handler(ctx, c pkgCqrs.Command) error {

	return nil
}

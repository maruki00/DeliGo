package handler

import (
	"context"
	"deligo/internal/product/domian/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type UpdateProductHandler struct {
	repo contract.IProductRepository
}

func NewUpdateProductHandler(repo contract.IProductRepository) UpdateProductHandler {
	return UpdateProductHandler{
		repo: repo,
	}
}

func (_this *UpdateProductHandler) handler(ctx context.Context, c pkgCqrs.Command) error {

	return nil
}

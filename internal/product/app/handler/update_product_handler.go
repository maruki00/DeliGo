package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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

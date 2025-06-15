package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/product/app/command"
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type DeleteProductHandler struct {
	repo contract.IProductRepository
}

func NewDeleteProductHandler(repo contract.IProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{
		repo: repo,
	}
}

func (_this *DeleteProductHandler) handle(ctx context.Context, c pkgCqrs.Command) error {

	cmd := c.(*command.DeleteProductCommand)
	return _this.repo.Delete(ctx, cmd.ID)
	return nil
}

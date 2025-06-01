package handlers

import (
	"context"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type DeleteShopHandler struct {
	repo contracts.IShopRepository
}

func NewDeleteShopHandler(repo contracts.IShopRepository) *DeleteShopHandler {
	return &DeleteShopHandler{
		repo: repo,
	}
}

func (_this *DeleteShopHandler) handler(ctx context.Context, command pkgCqrs.Command) error {

	cmd := command.(*commands.DeleteShopCommand)
	return _this.repo.Delete(ctx, cmd.ID)
}

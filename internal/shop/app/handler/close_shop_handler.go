package handler

import (
	"context"
	"deligo/internal/shop/app/command"
	"deligo/internal/shop/domain/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

const CLOSE = false

type CloseShopHandler struct {
	repo contract.IShopRepository
}

func NewCloseShopHandler(repo contract.IShopRepository) *CloseShopHandler {
	return &CloseShopHandler{
		repo: repo,
	}
}

func (_this *CloseShopHandler) handler(ctx context.Context, cmd pkgCqrs.Command) error {

	c := cmd.(*command.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, c.ID, CLOSE)
}

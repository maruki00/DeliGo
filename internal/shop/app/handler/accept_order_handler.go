package handler

import (
	"context"
	"deligo/internal/shop/app/command"
	"deligo/internal/shop/domain/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type AcceptOrderHandler struct {
	repo contract.IShopRepository
}

func NewAcceptOrderHandler(repo contract.IShopRepository) *AcceptOrderHandler {
	return &AcceptOrderHandler{
		repo: repo,
	}
}

func (_this *AcceptOrderHandler) handler(ctx context.Context, c pkgCqrs.Command) error {
	cmd := c.(*command.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, cmd.ID)
}

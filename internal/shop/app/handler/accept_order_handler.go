package handler

import (
	"context"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type AcceptOrderHandler struct {
	repo contracts.IShopRepository
}

func NewAcceptOrderHandler(repo contracts.IShopRepository) *AcceptOrderHandler {
	return &AcceptOrderHandler{
		repo: repo,
	}
}

func (_this *AcceptOrderHandler) handler(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, cmd.ID)
}

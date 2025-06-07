package handler

import (
	"context"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type UpdateShopHandler struct {
	repo contracts.IShopRepository
}

func NewUpdateShopHandler(repo contracts.IShopRepository) *UpdateShopHandler {
	return &UpdateShopHandler{
		repo: repo,
	}
}

func (_this *UpdateShopHandler) handler(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.UpdateShopCommand)
	return _this.repo.Update(ctx, cmd.ID, map[string]any{
		"name":     cmd.ShopName,
		"open_at":  cmd.OpenAt,
		"close_at": cmd.CloseAt,
	})
}

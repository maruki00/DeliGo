package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/shop/app/command"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateShopHandler struct {
	repo contract.IShopRepository
}

func NewUpdateShopHandler(repo contract.IShopRepository) *UpdateShopHandler {
	return &UpdateShopHandler{
		repo: repo,
	}
}

func (_this *UpdateShopHandler) handle(ctx context.Context, cmd pkgCqrs.Command) error {
	c := cmd.(*command.UpdateShopCommand)
	return _this.repo.Update(ctx, c.ID, map[string]any{
		"name":     c.ShopName,
		"open_at":  c.OpenAt,
		"close_at": c.CloseAt,
	})
}

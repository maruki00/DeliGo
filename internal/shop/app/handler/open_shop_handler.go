package handlers

import (
	"context"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type OpenShopHandler struct {
	repo contracts.IShopRepository
}

func NewOpenShopHandler(repo contracts.IShopRepository) *OpenShopHandler {
	return &OpenShopHandler{
		repo: repo,
	}
}

func (_this *OpenShopHandler) handler(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.OpenShopCommand)
	return _this.repo.UpdateStatus(ctx, cmd.ID, true)
}

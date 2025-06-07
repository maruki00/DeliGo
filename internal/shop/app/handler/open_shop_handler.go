package handler

import (
	"context"
	"deligo/internal/shop/app/command"
	"deligo/internal/shop/domain/contract"
	pkgCqrs "deligo/pkg/cqrs"
)

type OpenShopHandler struct {
	repo contract.IShopRepository
}

func NewOpenShopHandler(repo contract.IShopRepository) *OpenShopHandler {
	return &OpenShopHandler{
		repo: repo,
	}
}

func (_this *OpenShopHandler) handler(ctx context.Context, cmd pkgCqrs.Command) error {
	c := cmd.(*command.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, c.ID)
}

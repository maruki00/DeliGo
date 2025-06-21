package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/shop/app/command"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

const OPEN = true

type OpenShopHandler struct {
	repo contract.IShopRepository
}

func NewOpenShopHandler(repo contract.IShopRepository) *OpenShopHandler {
	return &OpenShopHandler{
		repo: repo,
	}
}

func (_this *OpenShopHandler) handle(ctx context.Context, cmd pkgCqrs.Command) error {
	c := cmd.(*command.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, c.ID, OPEN)
}

package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/shop/app/command"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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

func (_this *CloseShopHandler) handle(ctx context.Context, cmd pkgCqrs.Command) error {

	c := cmd.(*command.CloseShopCommand)
	return _this.repo.UpdateStatus(ctx, c.ID, CLOSE)
}

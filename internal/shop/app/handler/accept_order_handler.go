package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/shop/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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
	// cmd := c.(*command.AcceptOrderCommand)
	return nil
}

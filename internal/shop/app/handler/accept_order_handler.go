package handler

import (
	"context"
	"fmt"

	"github.com/maruki00/deligo/internal/shop/app/command"
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
	//TODO:
	// Call Accept Order By Grpc to accept the order
	cmd := c.(*command.AcceptOrderCommand)
	fmt.Println("Accept Order with id ", cmd.ID)
	return nil
}

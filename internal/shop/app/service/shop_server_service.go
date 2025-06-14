package service

import (
	"context"
	"errors"
	"time"

	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	"github.com/maruki00/deligo/internal/shop/app/command"
	shop_grpc "github.com/maruki00/deligo/internal/shop/infra/grpc/shop"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"

	"google.golang.org/grpc"
)

type ShopServerService struct {
	cmdBus *pkgCqrs.CommandBus
	qryBus *pkgCqrs.QueryBus
}

func NewShopServerService(cmdBus *pkgCqrs.CommandBus, qryBus *pkgCqrs.QueryBus) *ShopServerService {
	srv := ShopServerService{
		cmdBus: cmdBus,
		qryBus: qryBus,
	}
	return &srv
}

func (_this *ShopServerService) Save(ctx context.Context, in *shop_grpc.CreateShopRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) Update(ctx context.Context, in *shop_grpc.UpdateShopRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {

	id := sharedvo.Parse(in.ID)
	if id == "" {
		return nil, errors.New("invalid id")
	}

	openAt, err := time.Parse("2025-05-14", in.OpenAt)
	if err != nil {
		return nil, errors.New("invalid time on open_at")
	}

	closeAt, err := time.Parse("2025-05-14", in.CloseAt)
	if err != nil {
		return nil, errors.New("invalid time on close_at")
	}

	cmd := command.UpdateShopCommand{
		ID:       id,
		ShopName: in.ShopName,
		OpenAt:   openAt,
		CloseAt:  closeAt,
	}

	err = _this.cmdBus.Dispatch(ctx, &cmd)
	if err != nil {
		return nil, err
	}
	return &shop_grpc.ShopResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (_this *ShopServerService) Delete(ctx context.Context, in *shop_grpc.UpdateShopStatusRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	id := sharedvo.Parse(in.ID)
	if id == "" {
		return nil, errors.New("invalid id")
	}
	if err := _this.cmdBus.Dispatch(ctx, cmd) 

	return nil, nil
}

func (_this *ShopServerService) GetShopStatus(ctx context.Context, in *shop_grpc.GETRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) GetShop(ctx context.Context, in *shop_grpc.GETRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	return nil, nil
}

package service

import (
	"context"
	shop_grpc "deligo/internal/shop/infra/grpc/shop"
	pkgCqrs "deligo/pkg/cqrs"

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
	return nil, nil
}

func (_this *ShopServerService) Delete(ctx context.Context, in *shop_grpc.UpdateShopStatusRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) GetShopStatus(ctx context.Context, in *shop_grpc.GETRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) GetShop(ctx context.Context, in *shop_grpc.GETRequest, opts ...grpc.CallOption) (*shop_grpc.ShopResponse, error) {

	return nil, nil
}

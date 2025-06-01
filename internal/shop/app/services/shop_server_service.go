package services

import (
	"context"
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

func (_this *ShopServerService) Save(ctx context.Context, in *shop_grpc.CreateShopRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) Update(ctx context.Context, in *UpdateShopRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) Delete(ctx context.Context, in *UpdateShopStatusRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	return nil, nil
}

func (_this *ShopServerService) GetShopStatus(ctx context.Context, in *GETRequest, opts ...grpc.CallOption) (*ShopResponse, error) 	
return nil, nil
}

func (_this *ShopServerService) GetShop(ctx context.Context, in *GETRequest, opts ...grpc.CallOption) (*ShopResponse, error) {

	return nil, nil
}

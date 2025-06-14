package service

import (
	"context"

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

var response = &shop_grpc.ShopResponse{}

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

	cmd := command.UpdateShopCommand{
		ID:       sharedvo.NewID(in.ID),
		ShopName: in.ShopName,
		OpenAt:   in.OpenAt,
		CloseAt:  in.CloseAt,
	}

	return response, nil
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

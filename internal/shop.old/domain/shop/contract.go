package shop

import (
	"context"
)

// type ShopRepo interface {
// 	Save(context.Context, entity.ShopEntity)
// 	Delete(context.Context, vo.ID) error
// 	Update(context.Context, vo.ID, map[string]any) error
// 	UpdateStatus(context.Context, vo.ID, bool) error
// 	GetByID(context.Context, vo.ID) (*model.Shop, error)
// 	GetShopStatus(context.Context, vo.ID) (bool, error)
// }

type Repository interface {
	Create(ctx context.Context, s *Shop) error
	GetByID(ctx context.Context, id string) (*Shop, error)
	List(ctx context.Context, limit, offset int) ([]*Shop, error)
	Update(ctx context.Context, s *Shop) error
	Delete(ctx context.Context, id string) error
}

type Service interface {
	CreateShop(ctx context.Context, name, address, phone, ownerID string) (*Shop, error)
	GetShop(ctx context.Context, id string) (*Shop, error)
	ListShops(ctx context.Context, limit, offset int) ([]*Shop, error)
	UpdateShop(ctx context.Context, id, name, address, phone string) (*Shop, error)
	DeleteShop(ctx context.Context, id string) error
}

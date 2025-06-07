package contract

import (
	"context"
	sharedvo "deligo/internal/shared/valueobject"
	"deligo/internal/shop/domain/entity"
	"deligo/internal/shop/infra/model"
)

type IShopRepository interface {
	Save(ctx context.Context, shop entity.ShopEntity)
	Delete(ctx context.Context, id sharedvo.ID) error
	Update(ctx context.Context, id sharedvo.ID, fields map[string]any) error
	UpdateStatus(ctx context.Context, id sharedvo.ID, status bool) error
	GetByID(ctx context.Context, id sharedvo.ID) (*model.Shop, error)
	GetShopStatus(ctx context.Context, id sharedvo.ID) (bool, error)
}

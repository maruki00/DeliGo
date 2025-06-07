package contract

import (
	"context"
	sharedvo "deligo/internal/shared/valueobject"
	"deligo/internal/shop/domain/entity"
	"deligo/internal/shop/infra/model"
)

type IShopRepository interface {
	Save(context.Context, entity.ShopEntity)
	Delete(context.Context, sharedvo.ID) error
	Update(context.Context, sharedvo.ID, map[string]any) error
	UpdateStatus(context.Context, sharedvo.ID, bool) error
	GetByID(context.Context, sharedvo.ID) (*model.Shop, error)
	GetShopStatus(context.Context, sharedvo.ID) (bool, error)
}

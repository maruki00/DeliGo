package contracts

import (
	"context"
	shared_valueobject "deligo/internal/shared/value_objects"
	"deligo/internal/shop/domain/entity"
	"deligo/internal/shop/infra/model"
)

type IShopRepository interface {
	Save(context.Context, entity.ShopEntity) error
	Delete(context.Context, shared_valueobject.ID) error
	Update(context.Context, shared_valueobject.ID, map[string]any) error
	UpdateStatus(context.Context, shared_valueobject.ID, bool) error
	GetByID(context.Context, shared_valueobject.ID) (*model.Shop, error)
	GetShopStatus(context.Context, shared_valueobject.ID) (bool, error)
}

package contracts

import (
	"context"
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"deligo/internal/shop/domain/entities"
	"deligo/internal/shop/infra/models"
)

type IShopRepository interface {
	Save(context.Context, entities.ShopEntity) error
	Delete(context.Context, shared_valueobject.ID) error
	Update(context.Context, shared_valueobject.ID, map[string]any) error
	UpdateStatus(context.Context, shared_valueobject.ID, bool) error
	GetByID(context.Context, shared_valueobject.ID) (*models.Shop, error)
	GetShopStatus(context.Context, shared_valueobject.ID) (bool, error)
}

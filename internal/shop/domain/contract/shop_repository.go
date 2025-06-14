package contract

import (
	"context"

	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	"github.com/maruki00/deligo/internal/shop/domain/entity"
	"github.com/maruki00/deligo/internal/shop/infra/model"
)

type IShopRepository interface {
	Save(context.Context, entity.ShopEntity)
	Delete(context.Context, sharedvo.ID) error
	Update(context.Context, sharedvo.ID, map[string]any) error
	UpdateStatus(context.Context, sharedvo.ID, bool) error
	GetByID(context.Context, sharedvo.ID) (*model.Shop, error)
	GetShopStatus(context.Context, sharedvo.ID) (bool, error)
}

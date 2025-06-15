package contract

import (
	"context"

	"github.com/maruki00/deligo/internal/product/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type IProductRepository interface {
	Save(context.Context, *model.Product) error
	GetById(context.Context, sharedvo.ID) (*model.Product, error)
	List(context.Context, string) ([]*model.Product, error)
	Update(context.Context, sharedvo.ID, *model.Product) error
	Delete(context.Context, sharedvo.ID) error
	GetManyProductsByID(context.Context, []sharedvo.ID) ([]*model.Product, error)
	GetProductByMultipleId(context.Context, []sharedvo.ID) ([]*model.Product, error)
}

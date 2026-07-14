package contract

import (
	"context"

	"github.com/maruki00/deligo/internal/product/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type IProductRepository interface {
	Save(context.Context, *model.Product) error
	GetById(context.Context, sharedvo.ID) (*model.Product, error)
	Search(context.Context, string) ([]*model.Product, error)
	Update(context.Context, sharedvo.ID, *model.Product) error
	Delete(context.Context, sharedvo.ID) error
	GetMany(context.Context, []sharedvo.ID) ([]*model.Product, error)
	GetByIds(context.Context, []sharedvo.ID) ([]*model.Product, error)
}

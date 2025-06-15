package contract

import (
	"context"

	"github.com/maruki00/deligo/internal/product/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type IProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	GetById(ctx context.Context, id sharedvo.ID) (*model.Product, error)
	List(ctx context.Context, seasrch string) ([]*model.Product, error)
	Update(ctx context.Context, id sharedvo.ID, product *model.Product) error
	Delete(ctx context.Context, id sharedvo.ID) error
	GetManyProductsByID(ctx context.Context, ids []sharedvo.ID) ([]*model.Product, error)
	GetProductByMultipleId(ctx context.Context, ids []sharedvo.ID) ([]*model.Product, error)
}

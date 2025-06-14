package contract

import (
	"context"
	"github.com/maruki00/deligo/internal/product/infra/model"
)

type IProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	GetById(ctx context.Context, id string) (*model.Product, error)
	List(ctx context.Context, seasrch string) ([]*model.Product, error)
	Update(ctx context.Context, id int, product *model.Product) error
	Delete(ctx context.Context, id int) error
	GetManyProductsByID(ctx context.Context, ids []string) ([]*model.Product, error)
	GetProductByMultipleId(ctx context.Context, ids []string) ([]*model.Product, error)
}

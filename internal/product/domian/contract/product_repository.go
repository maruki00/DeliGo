package contract

import (
	"context"
	"deligo/internal/product/domian/entity"
)

type IProductRepository interface {
	Save(context.Context, entity.ProductEntity) error
	Find(context.Context, string) (entity.ProductEntity, error)
	Update(context.Context, string, map[string]any) error
	Delete(context.Context, string) error
	List(context.Context) ([]entity.ProductEntity, error)
	GetProductByMultipleId(context.Context, []string) ([]entity.ProductEntity, error)
}

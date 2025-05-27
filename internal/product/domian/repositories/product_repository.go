package repositories

import (
	"context"
	"deligo/internal/product/domian/entities"
)

type IProductRepository interface {
	Save(context.Context, entities.ProductEntity) error
	Find(context.Context, string) (entities.ProductEntity, error)
	Update(context.Context, string, map[string]any) error
	Delete(context.Context, string) error
	List(context.Context) ([]entities.ProductEntity, error)
	GetProductByMultipleId(context.Context, []string) ([]entities.ProductEntity, error)
}

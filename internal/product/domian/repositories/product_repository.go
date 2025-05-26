package repositories

import (
	"context"
	"deligo/internal/product/domian/entities"
)

type IProductRepository interface {
	Save(ctx context.Context, product entities.ProductEntity) (entities.ProductEntity, error)
	Find(ctx context.Context, id string) (entities.ProductEntity, error)
	Update(ctx context.Context, id string, data map[string]interface{}) (entities.ProductEntity, error)
	Delete(ctx context.Context, id string) (entities.ProductEntity, error)
	List(ctx context.Context) ([]entities.ProductEntity, error)
	// GetProductByMultipleId(ctx context.Context, ids []int) ([]models.Product, error)
}

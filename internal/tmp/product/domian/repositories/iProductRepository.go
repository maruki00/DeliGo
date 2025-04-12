package repositories

import (
	"context"
	"delivery/internal/product/domian/entities"
	"delivery/internal/product/infrastructure/models"
)

type IProductRepository interface {
	Insert(ctx context.Context, product entities.ProductEntity) (entities.ProductEntity, error)
	GetById(ctx context.Context, id int) (entities.ProductEntity, error)
	Search(ctx context.Context, seasrch string) ([]models.Product, error)
	Update(ctx context.Context, id int, data map[string]interface{}) (entities.ProductEntity, error)
	Delete(ctx context.Context, id int) (entities.ProductEntity, error)
	GetProductByMultipleId(ctx context.Context, ids []int) ([]models.Product, error)
}

package product

import "context"

type Repository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	ListByShop(ctx context.Context, shopID string, limit, offset int) ([]*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id string) error
}

type Service interface {
	CreateProduct(ctx context.Context, shopID, name, description string, price float64, stock int) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	ListProductsByShop(ctx context.Context, shopID string, limit, offset int) ([]*Product, error)
	UpdateProduct(ctx context.Context, id, name, description string, price float64) (*Product, error)
	AdjustStock(ctx context.Context, id string, delta int) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

package menu

import "context"

type Repository interface {
	Create(ctx context.Context, m *Menu) error
	GetByID(ctx context.Context, id string) (*Menu, error)
	ListByShop(ctx context.Context, shopID string) ([]*Menu, error)
	Update(ctx context.Context, m *Menu) error
	Delete(ctx context.Context, id string) error
}

type Service interface {
	CreateMenu(ctx context.Context, shopID, name string) (*Menu, error)
	GetMenu(ctx context.Context, id string) (*Menu, error)
	ListMenusByShop(ctx context.Context, shopID string) ([]*Menu, error)
	AddItem(ctx context.Context, menuID, productID string, price float64) (*Menu, error)
	RemoveItem(ctx context.Context, menuID, productID string) (*Menu, error)
	DeleteMenu(ctx context.Context, id string) error
}

package postgres

import (
	"context"

	menuDomain "github.com/maruki00/deligo/internal/shop/domain/menu"
	menuModel "github.com/maruki00/deligo/internal/shop/infra/model"
	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db: db}
}

func (r *MenuRepo) Create(ctx context.Context, m *menuDomain.Menu) error {
	items := make([]menuModel.Item, 0, len(m.Items))
	for _, item := range m.Items {
		items = append(items, menuModel.Item{ProductID: item.ProductID, Price: item.Price})
	}
	model := &menuModel.Menu{ID: m.ID, ShopID: m.ShopID, Name: m.Name, Items: items}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *MenuRepo) GetByID(ctx context.Context, id string) (*menuDomain.Menu, error) {
	var menu menuModel.Menu
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		return nil, err
	}
	items := make([]menuDomain.Item, 0, len(menu.Items))
	for _, item := range menu.Items {
		items = append(items, menuDomain.Item{ProductID: item.ProductID, Price: item.Price})
	}
	return &menuDomain.Menu{ID: menu.ID, ShopID: menu.ShopID, Name: menu.Name, Items: items}, nil
}

func (r *MenuRepo) ListByShop(ctx context.Context, shopID string) ([]*menuDomain.Menu, error) {
	var menus []*menuModel.Menu
	if err := r.db.WithContext(ctx).Where("shop_id = ?", shopID).Order("created_at desc").Find(&menus).Error; err != nil {
		return nil, err
	}

	items := make([]*menuDomain.Menu, 0, len(menus))
	for _, menu := range menus {
		converted := make([]menuDomain.Item, 0, len(menu.Items))
		for _, item := range menu.Items {
			converted = append(converted, menuDomain.Item{ProductID: item.ProductID, Price: item.Price})
		}
		items = append(items, &menuDomain.Menu{ID: menu.ID, ShopID: menu.ShopID, Name: menu.Name, Items: converted})
	}
	return items, nil
}

func (r *MenuRepo) Update(ctx context.Context, m *menuDomain.Menu) error {
	items := make([]menuModel.Item, 0, len(m.Items))
	for _, item := range m.Items {
		items = append(items, menuModel.Item{ProductID: item.ProductID, Price: item.Price})
	}
	model := &menuModel.Menu{ID: m.ID, ShopID: m.ShopID, Name: m.Name, Items: items}
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *MenuRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&menuModel.Menu{}, "id = ?", id).Error
}

var _ menuDomain.Repository = (*MenuRepo)(nil)

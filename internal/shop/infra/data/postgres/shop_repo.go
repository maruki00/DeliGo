package postgres

import (
	"context"

	shopDomain "github.com/maruki00/deligo/internal/shop/domain/shop"
	shopModel "github.com/maruki00/deligo/internal/shop/infra/model"
	"gorm.io/gorm"
)

type ShopRepo struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepo {
	return &ShopRepo{db: db}
}

func (r *ShopRepo) Create(ctx context.Context, s *shopDomain.Shop) error {
	model := &shopModel.Shop{ID: s.ID, Name: s.Name, Address: s.Address, Phone: s.Phone, OwnerID: s.OwnerID}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *ShopRepo) GetByID(ctx context.Context, id string) (*shopDomain.Shop, error) {
	var shop shopModel.Shop
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shopDomain.Shop{ID: shop.ID, Name: shop.Name, Address: shop.Address, Phone: shop.Phone, OwnerID: shop.OwnerID}, nil
}

func (r *ShopRepo) List(ctx context.Context, limit, offset int) ([]*shopDomain.Shop, error) {
	var shops []*shopModel.Shop
	query := r.db.WithContext(ctx).Order("created_at desc")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&shops).Error; err != nil {
		return nil, err
	}

	items := make([]*shopDomain.Shop, 0, len(shops))
	for _, shop := range shops {
		items = append(items, &shopDomain.Shop{ID: shop.ID, Name: shop.Name, Address: shop.Address, Phone: shop.Phone, OwnerID: shop.OwnerID})
	}
	return items, nil
}

func (r *ShopRepo) Update(ctx context.Context, s *shopDomain.Shop) error {
	model := &shopModel.Shop{ID: s.ID, Name: s.Name, Address: s.Address, Phone: s.Phone, OwnerID: s.OwnerID}
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *ShopRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&shopModel.Shop{}, "id = ?", id).Error
}

var _ shopDomain.Repository = (*ShopRepo)(nil)

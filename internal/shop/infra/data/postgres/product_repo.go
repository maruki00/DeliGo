package postgres

import (
	"context"

	productDomain "github.com/maruki00/deligo/internal/shop/domain/product"
	productModel "github.com/maruki00/deligo/internal/shop/infra/model"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(ctx context.Context, p *productDomain.Product) error {
	model := &productModel.Product{ID: p.ID, ShopID: p.ShopID, Name: p.Name, Description: p.Description, Price: p.Price, Stock: p.Stock}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *ProductRepo) GetByID(ctx context.Context, id string) (*productDomain.Product, error) {
	var product productModel.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &productDomain.Product{ID: product.ID, ShopID: product.ShopID, Name: product.Name, Description: product.Description, Price: product.Price, Stock: product.Stock}, nil
}

func (r *ProductRepo) ListByShop(ctx context.Context, shopID string, limit, offset int) ([]*productDomain.Product, error) {
	var products []*productModel.Product
	query := r.db.WithContext(ctx).Where("shop_id = ?", shopID).Order("created_at desc")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	items := make([]*productDomain.Product, 0, len(products))
	for _, product := range products {
		items = append(items, &productDomain.Product{ID: product.ID, ShopID: product.ShopID, Name: product.Name, Description: product.Description, Price: product.Price, Stock: product.Stock})
	}
	return items, nil
}

func (r *ProductRepo) Update(ctx context.Context, p *productDomain.Product) error {
	model := &productModel.Product{ID: p.ID, ShopID: p.ShopID, Name: p.Name, Description: p.Description, Price: p.Price, Stock: p.Stock}
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *ProductRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&productModel.Product{}, "id = ?", id).Error
}

var _ productDomain.Repository = (*ProductRepo)(nil)

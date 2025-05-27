package repositories

import (
	"context"
	"deligo/internal/product/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db pkgPostgres.PGHandler
}

func NewProductRepository(db pkgPostgres.PGHandler) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (_this *ProductRepository) Insert(ctx context.Context, product *models.Product) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := _this.db.GetDB().Model(&models.Product{}).Create(product).Error; err != nil {
			return err
		}
		return nil
	})
}

func (_this *ProductRepository) GetById(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	if err := _this.db.GetDB().Where("where = ? ", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (_this *ProductRepository) List(ctx context.Context, seasrch string) ([]*models.Product, error) {
	var items []*models.Product
	if err := _this.db.GetDB().Model(&models.Product{}).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (_this *ProductRepository) Update(ctx context.Context, id int, product *models.Product) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ProductRepository) Delete(ctx context.Context, id int) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (_this *ProductRepository) GetProductByMultipleId(ctx context.Context, ids []string) ([]*models.Product, error) {
	var items []*models.Product
	if err := _this.db.GetDB().Model(&models.Product{}).Where("id in ? ", ids).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil

}

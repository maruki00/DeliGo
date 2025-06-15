package repository

import (
	"context"

	"github.com/maruki00/deligo/internal/product/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	"github.com/maruki00/deligo/pkg/pagination"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

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

// Commands
func (_this *ProductRepository) Save(ctx context.Context, product *model.Product) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := _this.db.GetDB().Model(&model.Product{}).Create(product).Error; err != nil {
			return err
		}
		return nil
	})
}
func (_this *ProductRepository) Update(ctx context.Context, id sharedvo.ID, product *model.Product) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ProductRepository) Delete(ctx context.Context, id sharedvo.ID) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Product{}, "id = ?", id).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

// 	Search(context.Context, string) ([]*model.Product, error)
// 	GetMany(context.Context, []sharedvo.ID) ([]*model.Product, error)
// 	GetByIds(context.Context, []sharedvo.ID) ([]*model.Product, error)

// Queries
func (_this *ProductRepository) GetById(ctx context.Context, id sharedvo.ID) (*model.Product, error) {
	var product model.Product
	if err := _this.db.GetDB().Where(" id = ? ", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (_this *ProductRepository) Search(ctx context.Context, search string, pagination pagination.Pagination) ([]*model.Product, error) {
	var items []*model.Product
	if err := _this.db.GetDB().Model(&model.Product{}).Limit().Offset().Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (_this *ProductRepository) GetManyProductsByID(ctx context.Context, ids []sharedvo.ID) ([]*model.Product, error) {
	var items []*model.Product
	if err := _this.db.GetDB().Model(&model.Product{}).Where("id in ? ", ids).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

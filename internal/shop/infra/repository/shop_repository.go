package repository

import (
	"context"

	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	"github.com/maruki00/deligo/internal/shop/domain/entity"
	"github.com/maruki00/deligo/internal/shop/infra/model"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

	"gorm.io/gorm"
)

type ShopRepository struct {
	db *pkgPostgres.PGHandler
}

func NewShopRepository(db *pkgPostgres.PGHandler) *ShopRepository {
	return &ShopRepository{
		db: db,
	}
}

func (_this *ShopRepository) Save(ctx context.Context, shop entity.ShopEntity) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Shop{}).Create(&shop).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) Delete(ctx context.Context, id sharedvo.ID) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? ", id).Delete(&model.Shop{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) Update(ctx context.Context, id sharedvo.ID, fields map[string]any) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Shop{}).Where("id = ?", id).UpdateColumns(fields).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) UpdateStatus(ctx context.Context, id sharedvo.ID, status bool) error {
	res := _this.db.GetDB().Model(model.Shop{}).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (_this *ShopRepository) GetByID(ctx context.Context, id sharedvo.ID) (*model.Shop, error) {
	var shop model.Shop
	if err := _this.db.GetDB().Model(&model.Shop{}).Where("id = ?", id).Find(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (_this *ShopRepository) GetShopStatus(ctx context.Context, id sharedvo.ID) (bool, error) {
	var status bool
	if err := _this.db.GetDB().Model(&model.Shop{}).Select("status").Where("id = ?", id).Scan(&status).Error; err != nil {
		return false, err
	}
	return true, nil
}

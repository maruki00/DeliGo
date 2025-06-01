package respositories

import (
	"context"
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"deligo/internal/shop/domain/entities"
	"deligo/internal/shop/infra/models"
	pkgPostgres "deligo/pkg/postgres"

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

func (_this *ShopRepository) Save(ctx context.Context, shop entities.ShopEntity) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Shop{}).Create(&shop).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) Delete(ctx context.Context, id shared_valueobject.ID) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? ", id).Delete(&models.Shop{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) Update(ctx context.Context, id shared_valueobject.ID, fields map[string]any) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Shop{}).Where("id = ?", id).UpdateColumns(fields).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}

func (_this *ShopRepository) UpdateStatus(ctx context.Context, id shared_valueobject.ID, status bool) error {
	res := _this.db.GetDB().Model(models.Shop{}).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (_this *ShopRepository) GetByID(ctx context.Context, id shared_valueobject.ID) (*models.Shop, error) {
	var shop models.Shop
	if err := _this.db.GetDB().Model(&models.Shop{}).Where("id = ?", id).Find(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (_this *ShopRepository) GetShopStatus(ctx context.Context, id shared_valueobject.ID) (bool, error) {
	var status bool
	if err := _this.db.GetDB().Model(&models.Shop{}).Select("status").Where("id = ?", id).Scan(&status).Error; err != nil {
		return false, err
	}
	return true, nil
}

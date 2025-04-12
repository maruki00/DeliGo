package respositories

import (
	"delivery/internal/shop/domain/entities"

	"gorm.io/gorm"
)

type ShopRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewShopRepository(db *gorm.DB, model interface{}) *ShopRepository {
	return &ShopRepository{
		db:    db,
		model: model,
	}
}

func (obj *ShopRepository) SetStatus(id int, status int) error {

	res := obj.db.Model(obj.model).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (obj *ShopRepository) UpdateShopSettings(id int, shop entities.ShopEntity) (entities.ShopEntity, error) {
	res := obj.db.Model(obj.model).Where("id = ? ", id).Updates(shop)
	if res.Error != nil {
		return nil, res.Error
	}
	return shop, nil
}

package respositories

import (
	"delivery/internal/shop/domain/entities"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db    *gorm.Model
	model interface{}
}

func NewMenuRepository(db *gorm.Model, model interface{}) *MenuRepository {
	return &MenuRepository{
		db:    db,
		model: model,
	}
}

func (obj *MenuRepository) Make(menn entities.MenuEntity) (entities.MenuEntity, error) {

	// res := obj.

	return nil, nil
}

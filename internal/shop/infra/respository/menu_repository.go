package repository

import (
	"deligo/internal/shop/domain/entity"

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

func (obj *MenuRepository) Make(menn entity.MenuEntity) (entity.MenuEntity, error) {

	// res := obj.

	return nil, nil
}

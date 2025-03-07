package models

import (
	"time"

	"gorm.io/gorm"
)

type MenuProduct struct {
	gorm.Model
	Id        int       `json: "id"`
	MenuId    int       `json: "menu_id"`
	ProductId int       `json: "product_id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func (obj *MenuProduct) GetId() int {
	return obj.Id
}
func (obj *MenuProduct) GetMenuId() int {
	return obj.MenuId
}
func (obj *MenuProduct) GetProductId() int {
	return obj.ProductId
}

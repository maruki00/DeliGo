package model

import shared_valueobject "github.com/maruki00/deligo/internal/shared/domain/sharedvo"

type MenuProduct struct {
	MenuId    shared_valueobject.ID `json: "menu_id"`
	ProductId shared_valueobject.ID `json: "product_id"`
}

func (obj *MenuProduct) GetMenuId() shared_valueobject.ID {
	return obj.MenuId
}
func (obj *MenuProduct) GetProductId() shared_valueobject.ID {
	return obj.ProductId
}

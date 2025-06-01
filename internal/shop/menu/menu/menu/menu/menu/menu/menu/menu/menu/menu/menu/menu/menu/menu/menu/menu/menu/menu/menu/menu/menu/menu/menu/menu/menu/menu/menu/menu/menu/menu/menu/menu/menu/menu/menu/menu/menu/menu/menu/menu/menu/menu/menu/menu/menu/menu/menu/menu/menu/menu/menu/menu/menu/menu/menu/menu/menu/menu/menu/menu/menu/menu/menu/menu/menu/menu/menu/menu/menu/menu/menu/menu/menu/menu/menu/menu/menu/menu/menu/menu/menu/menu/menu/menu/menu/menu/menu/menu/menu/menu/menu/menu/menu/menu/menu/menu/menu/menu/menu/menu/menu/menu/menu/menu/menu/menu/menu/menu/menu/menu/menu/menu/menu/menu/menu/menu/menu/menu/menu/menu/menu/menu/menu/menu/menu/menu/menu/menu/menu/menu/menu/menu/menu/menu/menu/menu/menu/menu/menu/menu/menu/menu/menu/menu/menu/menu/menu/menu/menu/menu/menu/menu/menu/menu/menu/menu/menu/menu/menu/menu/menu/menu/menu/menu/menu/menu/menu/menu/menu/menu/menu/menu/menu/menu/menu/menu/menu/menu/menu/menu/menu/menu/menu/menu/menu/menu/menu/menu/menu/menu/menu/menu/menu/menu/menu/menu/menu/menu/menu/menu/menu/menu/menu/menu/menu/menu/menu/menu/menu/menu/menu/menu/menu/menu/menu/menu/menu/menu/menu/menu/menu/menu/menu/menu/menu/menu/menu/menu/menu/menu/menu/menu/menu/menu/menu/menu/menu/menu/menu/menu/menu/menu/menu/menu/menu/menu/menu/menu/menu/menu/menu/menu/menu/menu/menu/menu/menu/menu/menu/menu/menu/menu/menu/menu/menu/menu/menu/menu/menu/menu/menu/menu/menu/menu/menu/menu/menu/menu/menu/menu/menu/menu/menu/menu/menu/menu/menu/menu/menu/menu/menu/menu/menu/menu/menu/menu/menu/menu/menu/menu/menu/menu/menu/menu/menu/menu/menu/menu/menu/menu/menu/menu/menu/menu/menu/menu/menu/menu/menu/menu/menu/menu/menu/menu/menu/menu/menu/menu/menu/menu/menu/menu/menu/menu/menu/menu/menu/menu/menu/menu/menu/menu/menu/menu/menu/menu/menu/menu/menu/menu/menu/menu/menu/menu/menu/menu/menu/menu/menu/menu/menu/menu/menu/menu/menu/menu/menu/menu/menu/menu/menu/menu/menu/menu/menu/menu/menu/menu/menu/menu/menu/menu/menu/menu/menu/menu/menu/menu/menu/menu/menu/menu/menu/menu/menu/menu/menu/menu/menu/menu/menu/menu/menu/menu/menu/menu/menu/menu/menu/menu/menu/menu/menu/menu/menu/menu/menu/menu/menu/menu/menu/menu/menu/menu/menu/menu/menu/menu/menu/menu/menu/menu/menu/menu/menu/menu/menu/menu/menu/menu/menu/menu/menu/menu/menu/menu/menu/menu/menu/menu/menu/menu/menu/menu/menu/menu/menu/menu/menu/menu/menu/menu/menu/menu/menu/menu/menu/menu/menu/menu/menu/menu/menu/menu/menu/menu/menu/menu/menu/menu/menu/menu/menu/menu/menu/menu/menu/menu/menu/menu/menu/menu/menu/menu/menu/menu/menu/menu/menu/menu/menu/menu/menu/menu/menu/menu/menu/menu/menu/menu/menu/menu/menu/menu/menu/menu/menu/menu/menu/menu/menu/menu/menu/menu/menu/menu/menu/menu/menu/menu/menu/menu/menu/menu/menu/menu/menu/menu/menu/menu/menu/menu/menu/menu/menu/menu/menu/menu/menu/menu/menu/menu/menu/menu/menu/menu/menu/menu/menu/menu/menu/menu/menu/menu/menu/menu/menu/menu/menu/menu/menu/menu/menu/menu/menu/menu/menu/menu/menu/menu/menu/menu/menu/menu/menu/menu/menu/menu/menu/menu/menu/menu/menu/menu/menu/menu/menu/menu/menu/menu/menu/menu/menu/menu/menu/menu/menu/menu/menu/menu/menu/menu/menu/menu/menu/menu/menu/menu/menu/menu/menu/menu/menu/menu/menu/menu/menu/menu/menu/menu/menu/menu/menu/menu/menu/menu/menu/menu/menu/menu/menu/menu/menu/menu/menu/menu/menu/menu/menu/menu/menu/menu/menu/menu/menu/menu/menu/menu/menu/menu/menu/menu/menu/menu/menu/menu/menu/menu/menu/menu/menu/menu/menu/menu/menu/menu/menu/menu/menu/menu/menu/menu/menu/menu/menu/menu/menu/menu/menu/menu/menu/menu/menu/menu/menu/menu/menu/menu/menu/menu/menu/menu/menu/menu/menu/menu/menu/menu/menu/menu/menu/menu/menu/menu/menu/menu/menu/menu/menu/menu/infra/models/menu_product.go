package models

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

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

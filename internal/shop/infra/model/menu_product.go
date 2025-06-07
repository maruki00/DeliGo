package model

import sharedvo "deligo/internal/shared/valueobject"

type MenuProduct struct {
	MenuId    sharedvo.ID `json: "menu_id"`
	ProductId sharedvo.ID `json: "product_id"`
}

func (obj *MenuProduct) GetMenuId() sharedvo.ID {
	return obj.MenuId
}
func (obj *MenuProduct) GetProductId() sharedvo.ID {
	return obj.ProductId
}

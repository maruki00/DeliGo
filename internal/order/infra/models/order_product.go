package models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	Id        int `json: "id"`
	OrderId   int `json: "order_id"`
	ProductId int `json: "product_id"`
	Qty       int `json: "qty"`
}

func (obj *OrderProduct) GetId() int {
	return obj.Id
}
func (obj *OrderProduct) GetOrderId() int {
	return obj.OrderId
}
func (obj *OrderProduct) GetProductId() int {
	return obj.ProductId
}
func (obj *OrderProduct) GetQty() int {
	return obj.Qty
}

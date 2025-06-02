package entities

type OrderProductEntity interface {
	GetOrderId() int
	GetProductId() int
	GetQty() int
	GetUnitPrice() float32
}

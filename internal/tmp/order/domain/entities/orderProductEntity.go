package entities

type OrderProductEntity interface {
	GetId() int
	GetOrderId() int
	GetProductId() int
	GetQty() int
}

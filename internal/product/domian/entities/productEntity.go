package entities

type ProductEntity interface {
	GetId() string
	GetLabel() string
	GetPrice() float32
	GetQty() int
}

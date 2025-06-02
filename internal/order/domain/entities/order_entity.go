package entities

import (
	valueobjects "deligo/internal/order/domain/value_objects"
	shared_valueobject "deligo/internal/shared/value_objects"
	"time"
)

type OrderEntity interface {
	CostumerId() int
	OrderStatus() int
	TotalAmount() float32
	Currency() valueobjects.Currency
	ShippingAddress() string
	BillingAddress() string
	OrderDate() time.Time
	PayementID() shared_valueobject.ID
}

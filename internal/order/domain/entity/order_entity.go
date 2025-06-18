package entity

import (
	"time"

	"github.com/maruki00/deligo/internal/order/domain/valueobject"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type OrderEntity interface {
	CostumerId() int
	OrderStatus() int
	TotalAmount() float32
	Currency() valueobject.Currency
	ShippingAddress() string
	BillingAddress() string
	OrderDate() time.Time
	PayementID() sharedvo.ID
}

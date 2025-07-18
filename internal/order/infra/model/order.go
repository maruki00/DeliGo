package model

import (
	"time"

	"github.com/maruki00/deligo/internal/order/domain/valueobject"
	shared_model "github.com/maruki00/deligo/internal/shared/model"
	shared_valueobject "github.com/maruki00/deligo/internal/shared/value_object"
)

type Order struct {
	shared_model.BaseModel
	CostumerId      int                   `json:"costumer_id"`
	OrderStatus     int                   `json:"order_status"`
	TotalAmount     float32               `json:"total_amount"`
	Currency        valueobject.Currency `json:"currency"`
	ShippingAddress string                `json:"shipping_address"`
	BillingAddress  string                `json:"billing_address"`
	OrderDate       time.Time             `json:"order_date"`
	PayementID      shared_valueobject.ID `json:"payement_id"`
}

func (_this *Order) GetCostumerId() int {
	return _this.CostumerId
}
func (_this *Order) GetOrderStatus() int {
	return _this.OrderStatus
}
func (_this *Order) GetTotalAmount() float32 {
	return _this.TotalAmount
}
func (_this *Order) GetCurrency() valueobject.Currency {
	return _this.Currency
}
func (_this *Order) GetShippingAddress() string {
	return _this.ShippingAddress
}
func (_this *Order) GetBillingAddress() string {
	return _this.BillingAddress
}
func (_this *Order) GetOrderDate() time.Time {
	return _this.OrderDate
}
func (_this *Order) GetPayementID() shared_valueobject.ID {
	return _this.PayementID
}

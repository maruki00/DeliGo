package models

import (
	valueobjects "deligo/internal/order/domain/value_objects"
	shared_models "deligo/internal/shared/infra/models"
	shared_valueobject "deligo/internal/shared/value_objects"
	"time"
)

type Order struct {
	shared_models.BaseModel
	CostumerId      int                   `json:"costumer_id"`
	OrderStatus     int                   `json:"order_status"`
	TotalAmount     float32               `json:"total_amount"`
	Currency        valueobjects.Currency `json:"currency"`
	ShippingAddress string                `json:"shipping_address"`
	BillingAddress  string                `json:"billing_address"`
	OrderDate       time.Time             `json:"order_date"`
	PayementID      shared_valueobject.ID `json:"payement_id"`
}

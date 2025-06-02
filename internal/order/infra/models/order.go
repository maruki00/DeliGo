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

// id VARCHAR(36) PRIMARY KEY NOT NULL,
//     customer_id VARCHAR(36) NOT NULL,
//     order_status VARCHAR(50) NOT NULL,
//     total_amount DECIMAL(10, 2) NOT NULL,
//     currency VARCHAR(3) NOT NULL,
//     order_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     shipping_address JSONB NULL,
//     billing_address JSONB NULL,
//     payment_id VARCHAR(36) NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

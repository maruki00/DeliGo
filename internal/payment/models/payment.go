package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	StatusInitiated PaymentStatus = "initiated"
	StatusSucceeded PaymentStatus = "succeeded"
	StatusFailed    PaymentStatus = "failed"
	StatusRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID             string        `gorm:"type:varchar(36);primaryKey" json:"id"`
	OrderID        string        `gorm:"type:varchar(36);not null;index:idx_order" json:"order_id"`
	CustomerID     string        `gorm:"type:varchar(36);not null" json:"customer_id"`
	StripeChargeID *string       `gorm:"type:varchar(255);unique" json:"stripe_charge_id"`
	Amount         float64       `gorm:"type:numeric(10,2);not null" json:"amount"`
	Currency       string        `gorm:"type:varchar(3);default:'USD'" json:"currency"`
	Status         PaymentStatus `gorm:"type:varchar(20);default:'initiated'" json:"status"`
	CreatedAt      time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

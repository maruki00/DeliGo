package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID               int       `json:"id"`
	OrderFingerprint string    `json:"order_fingerprint"`
	CostumerId       int       `json:"costumer_id"`
	Cost             float32   `json: "cost"`
	Status           int       `json: "status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (obj *Order) GetId() int {
	return obj.Id
}
func (obj *Order) GetOrderFingerprint() string {
	return obj.OrderFingerprint
}
func (obj *Order) GetCostumerId() int {
	return obj.CostumerId
}
func (obj *Order) GetCost() float32 {
	return obj.Cost
}
func (obj *Order) GetStatus() int {
	return obj.Status
}

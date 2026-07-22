package domain

import (
	"context"
	"time"
)

type OrderStatus string

const (
	StatusPendingPayment OrderStatus = "pending_payment"
	StatusPaid           OrderStatus = "paid"
	StatusPreparing      OrderStatus = "preparing"
	StatusReadyForPickup OrderStatus = "ready_for_pickup"
	StatusPickedUp       OrderStatus = "picked_up"
	StatusDelivered      OrderStatus = "delivered"
	StatusCancelled      OrderStatus = "cancelled"
)

type Order struct {
	ID              string      `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID      string      `gorm:"type:varchar(36);not null;index:idx_customer" json:"customer_id"`
	RestaurantID    string      `gorm:"type:varchar(36);not null;index:idx_restaurant" json:"restaurant_id"`
	Status          OrderStatus `gorm:"type:varchar(50);default:'pending_payment';not null" json:"status"`
	TotalAmount     float64     `gorm:"type:numeric(10,2);not null" json:"total_amount"`
	DeliveryAddress string      `gorm:"type:varchar(255);not null" json:"delivery_address"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"items"`
	CreatedAt       time.Time   `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time   `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type OrderItem struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     string  `gorm:"type:uuid;not null;index" json:"order_id"`
	ProductID   string  `gorm:"type:varchar(36);not null" json:"product_id"`
	ProductName string  `gorm:"type:varchar(255);not null" json:"product_name"`
	Price       float64 `gorm:"type:numeric(10,2);not null" json:"price"`
	Quantity    int     `gorm:"type:integer;not null" json:"quantity"`
}

type OrderEvent struct {
	EventName string    `json:"event_name"`
	OrderID   string    `json:"order_id"`
	Payload   Order     `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	FindByID(ctx context.Context, id string) (*Order, error)
	UpdateStatus(ctx context.Context, id string, currentStatus, newStatus OrderStatus) error
	Update(ctx context.Context, order *Order) error
}

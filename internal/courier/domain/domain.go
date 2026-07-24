package domain

import (
	"time"
)

type CourierService interface {
	HandleOrderConfirmed(orderID string, restaurantLat, restaurantLon float64) error
	AcceptOrder(orderID string, courierID string) error
	ArriveAtRestaurant(orderID string, courierID string) error
	StartCourier(orderID string, courierID string) error
	CompleteCourier(orderID string, courierID string) error
	TrackLocation(courierID string, orderID string, lat, lon float64) error
	RegisterCourier(id string, vehicle string, lat, lon float64, active bool) error
}

type CourierRepository interface {
	GetActiveCouriers() ([]Courier, error)
	GetCourierBusyStatus(courierID string) (bool, error)
	CreateOrderCourier(assignment *OrderCourier) error
	UpdateOrderCourier(assignment *OrderCourier) error
	GetOrderCourierByOrder(orderID string) (*OrderCourier, error)
	UpdateCourierLocation(courierID string, lat, lon float64) error
	SaveTrackingLog(log *OrderTracking) error
	CreateCourier(courier *Courier) error
}

type Courier struct {
	ID               string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	VehicleType      string    `gorm:"type:varchar(20);not null" json:"vehicle_type"`
	IsActive         bool      `gorm:"default:false" json:"is_active"`
	CurrentLatitude  float64   `gorm:"type:decimal(9,6)" json:"current_latitude"`
	CurrentLongitude float64   `gorm:"type:decimal(9,6)" json:"current_longitude"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrderCourier struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     string     `gorm:"type:varchar(36);not null;index" json:"order_id"`
	CourierID   *string    `gorm:"type:varchar(36);index" json:"courier_id"`
	Status      string     `gorm:"type:varchar(20);default:'searching'" json:"status"`
	AssignedAt  *time.Time `json:"assigned_at"`
	PickedUpAt  *time.Time `json:"picked_up_at"`
	DeliveredAt *time.Time `json:"delivered_at"`
}

type OrderTracking struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID    string    `gorm:"type:varchar(36);not null;index" json:"order_id"`
	Latitude   float64   `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Longitude  float64   `gorm:"type:decimal(9,6);not null" json:"longitude"`
	RecordedAt time.Time `gorm:"autoCreateTime" json:"recorded_at"`
}

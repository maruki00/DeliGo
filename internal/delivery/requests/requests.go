package requests

type CourierSetupRequest struct {
	ID          string  `json:"id" binding:"required"`
	VehicleType string  `json:"vehicle_type" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	IsActive    bool    `json:"is_active"`
}

type OrderStatusChangeRequest struct {
	CourierID string `json:"courier_id" binding:"required"`
}

type LocationPingRequest struct {
	CourierID string  `json:"courier_id" binding:"required"`
	OrderID   string  `json:"order_id"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

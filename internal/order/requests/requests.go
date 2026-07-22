package requests

type CreateOrderItemRequest struct {
	ProductID   string  `json:"product_id" binding:"required"`
	ProductName string  `json:"product_name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
}

type CreateOrderRequest struct {
	CustomerID      string                   `json:"customer_id" binding:"required"`
	RestaurantID    string                   `json:"restaurant_id" binding:"required"`
	DeliveryAddress string                   `json:"delivery_address" binding:"required"`
	Items           []CreateOrderItemRequest `json:"items" binding:"required,dive"`
}

type AcceptOrderRequest struct {
	Actor string `json:"actor" binding:"required,oneof=restaurant courier"`
}

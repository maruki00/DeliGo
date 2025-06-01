package requests

type OrderProduct struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type CreateOrderRequest struct {
	CostomerId    int            `json:"costomer_id"`
	OrderProducts []OrderProduct `json:"order_products"`
}

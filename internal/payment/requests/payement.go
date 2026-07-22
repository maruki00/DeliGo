package requests

type ChargeRequest struct {
	OrderID    string  `json:"order_id" binding:"required"`
	CustomerID string  `json:"customer_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required,gt=0"`
	Currency   string  `json:"currency" binding:"required,len=3"`
	Token      string  `json:"token" binding:"required"` // Stripe source/tok_ mock token
}

type RefundRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

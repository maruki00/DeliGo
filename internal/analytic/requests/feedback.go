package requests

type CreateFeedbackRequest struct {
	OrderID       string `json:"order_id" binding:"required,uuid4"`
	CustomerID    string `json:"customer_id" binding:"required,uuid4"`
	ProductID     string `json:"product_id" binding:"required,uuid4"`
	ProductRating int    `json:"product_rating" binding:"required,min=1,max=5"`
	ProductReview string `json:"product_review"`
	CourierRating int    `json:"courier_rating" binding:"required,min=1,max=5"`
	CourierReview string `json:"courier_review"`
	ReportedIssue string `json:"reported_issue"`
}

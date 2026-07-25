package domain

import (
	"time"

	"github.com/maruki00/deligo/internal/analytic/requests"
)

type AnalyticRepository interface {
	UpdateMetrics(productID string) error
	GetByProductID(productID string) (*Analytic, error)
}

type FeedbackRepository interface {
	Save(feedback *Feedback) error
}

type FeedbackService interface {
	SubmitFeedback(req *requests.CreateFeedbackRequest) (*Feedback, error)
	GetAnalyticsByProductID(productID string) (*Analytic, error)
}

type EventPublisher interface {
	PublishFeedbackCreated(feedback *Feedback) error
}

type Feedback struct {
	ID            string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	OrderID       string    `gorm:"type:varchar(36);not null;index:idx_order" json:"order_id"`
	CustomerID    string    `gorm:"type:varchar(36);not null" json:"customer_id"`
	ProductID     string    `gorm:"type:varchar(36);not null;index:idx_product" json:"product_id"` // Added dynamically to map analytics down to specific products
	ProductRating int       `gorm:"type:int;check:product_rating BETWEEN 1 AND 5" json:"product_rating"`
	ProductReview string    `gorm:"type:text" json:"product_review"`
	CourierRating int       `gorm:"type:int;check:courier_rating BETWEEN 1 AND 5" json:"courier_rating"`
	CourierReview string    `gorm:"type:text" json:"courier_review"`
	ReportedIssue string    `gorm:"type:text" json:"reported_issue"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Analytic struct {
	ID                   uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID            string    `gorm:"type:varchar(36);uniqueIndex" json:"product_id"`
	TotalReviews         int       `json:"total_reviews"`
	AverageProductRating float64   `json:"average_product_rating"`
	TotalIssuesReported  int       `json:"total_issues_reported"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// service/feedback_service.go
package service

import (
	"time"

	"github.com/maruki00/deligo/internal/analytic/domain"
	"github.com/maruki00/deligo/internal/analytic/requests"

	"github.com/google/uuid"
)

type FeedbackService struct {
	feedbackRepo domain.FeedbackRepository
	analyticRepo domain.AnalyticRepository
	publisher    domain.EventPublisher
}

func NewFeedbackService(fRepo domain.FeedbackRepository, aRepo domain.AnalyticRepository, pub domain.EventPublisher) *FeedbackService {
	return &FeedbackService{
		feedbackRepo: fRepo,
		analyticRepo: aRepo,
		publisher:    pub,
	}
}

func (s *FeedbackService) SubmitFeedback(req *requests.CreateFeedbackRequest) (*domain.Feedback, error) {
	fb := &domain.Feedback{
		ID:            uuid.New().String(),
		OrderID:       req.OrderID,
		CustomerID:    req.CustomerID,
		ProductID:     req.ProductID,
		ProductRating: req.ProductRating,
		ProductReview: req.ProductReview,
		CourierRating: req.CourierRating,
		CourierReview: req.CourierReview,
		ReportedIssue: req.ReportedIssue,
		CreatedAt:     time.Now(),
	}

	// 1. Synchronously commit the record to the core feedback domain store
	if err := s.feedbackRepo.Save(fb); err != nil {
		return nil, err
	}

	// 2. Fire-and-forget payload over message broker pipeline for fast execution tracking returns
	if err := s.publisher.PublishFeedbackCreated(fb); err != nil {
		// Production Note: In real production systems, apply an Outbox Pattern
		// if message loss must be structurally prevented if broker cuts out briefly.
	}

	return fb, nil
}

func (s *FeedbackService) GetAnalyticsByProductID(productID string) (*domain.Analytic, error) {
	return s.analyticRepo.GetByProductID(productID)
}

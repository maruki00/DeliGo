// service/feedback_service.go
package service

import (
	"time"

	"github.com/maruki00/deligo/internal/analytic/domain"
	"github.com/maruki00/deligo/internal/analytic/messaging"
	"github.com/maruki00/deligo/internal/analytic/repository"
	"github.com/maruki00/deligo/internal/analytic/requests"

	"github.com/google/uuid"
)

type FeedbackService interface {
	SubmitFeedback(req *requests.CreateFeedbackRequest) (*domain.Feedback, error)
	GetAnalyticsByProductID(productID string) (*domain.Analytic, error)
}

type feedbackService struct {
	feedbackRepo repository.FeedbackRepository
	analyticRepo repository.AnalyticRepository
	publisher    messaging.EventPublisher
}

func NewFeedbackService(fRepo repository.FeedbackRepository, aRepo repository.AnalyticRepository, pub messaging.EventPublisher) FeedbackService {
	return &feedbackService{
		feedbackRepo: fRepo,
		analyticRepo: aRepo,
		publisher:    pub,
	}
}

func (s *feedbackService) SubmitFeedback(req *requests.CreateFeedbackRequest) (*domain.Feedback, error) {
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

func (s *feedbackService) GetAnalyticsByProductID(productID string) (*domain.Analytic, error) {
	return s.analyticRepo.GetByProductID(productID)
}

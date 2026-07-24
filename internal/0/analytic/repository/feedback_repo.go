package repository

import (
	"github.com/maruki00/deligo/internal/analytic/domain"

	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Save(feedback *domain.Feedback) error
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) Save(feedback *domain.Feedback) error {
	return r.db.Create(feedback).Error
}

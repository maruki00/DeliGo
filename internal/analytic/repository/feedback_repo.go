package repository

import (
	"github.com/maruki00/deligo/internal/analytic/domain"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{db: db}
}

func (r *FeedbackRepository) Save(feedback *domain.Feedback) error {
	return r.db.Create(feedback).Error
}

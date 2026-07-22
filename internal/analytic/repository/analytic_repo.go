package repository

import (
	"time"

	"github.com/maruki00/deligo/internal/analytic/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnalyticRepository interface {
	UpdateMetrics(productID string) error
	GetByProductID(productID string) (*domain.Analytic, error)
}

type analyticRepository struct {
	db *gorm.DB
}

func NewAnalyticRepository(db *gorm.DB) AnalyticRepository {
	return &analyticRepository{db: db}
}

func (r *analyticRepository) GetByProductID(productID string) (*domain.Analytic, error) {
	var analytic domain.Analytic
	err := r.db.Where("product_id = ?", productID).First(&analytic).Error
	if err != nil {
		return nil, err
	}
	return &analytic, nil
}

func (r *analyticRepository) UpdateMetrics(productID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var summary struct {
			TotalReviews int
			AvgRating    float64
			TotalIssues  int
		}

		err := tx.Model(&domain.Feedback{}).
			Select("COUNT(id) as total_reviews, COALESCE(AVG(product_rating), 0) as avg_rating, SUM(CASE WHEN reported_issue IS NOT NULL AND reported_issue != '' THEN 1 ELSE 0 END) as total_issues").
			Where("product_id = ?", productID).
			Scan(&summary).Error

		if err != nil {
			return err
		}

		analytic := domain.Analytic{
			ProductID:            productID,
			TotalReviews:         summary.TotalReviews,
			AverageProductRating: summary.AvgRating,
			TotalIssuesReported:  summary.TotalIssues,
			UpdatedAt:            time.Now(),
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "product_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"total_reviews", "average_product_rating", "total_issues_reported", "updated_at"}),
		}).Create(&analytic).Error
	})
}

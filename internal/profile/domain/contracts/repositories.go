package contracts

import (
	"context"
	"deligo/internal/profile/infra/models"
)

type Pagination struct{}

type IPorofileRepository interface {
	Save(ctx context.Context, policy *models.Profile) error
	Disable(ctx context.Context, policy *models.Profile) error
	FindByUserID(ctx context.Context, id string) (*models.Policy, error)
	FindByName(ctx context.Context, name string) (*models.Policy, error)
	Delete(ctx context.Context, id string) error
	ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error)
}

package contracts

import (
	"context"
	"deligo/internal/profile/infra/models"
)

type Pagination struct{}

type IPorofileRepository interface {
	FindByID(ctx context.Context, id string) (*models.Policy, error)
	FindByName(ctx context.Context, name string) (*models.Policy, error)
	Save(ctx context.Context, policy *models.Policy) error
	Delete(ctx context.Context, id string) error
	ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error)
}

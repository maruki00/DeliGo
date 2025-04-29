package repositories

import (
	"context"
	"deligo/internal/iam/infra/models"
)

type PolicyRepository struct{}

func (_this *PolicyRepository) Save(ctx context.Context, policy *models.Policy) error {
	return nil
}

func (_this *PolicyRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (_this *PolicyRepository) FindByID(ctx context.Context, id string) (*models.Policy, error) {
	return nil, nil
}
func (_this *PolicyRepository) FindByName(ctx context.Context, name string) (*models.Policy, error) {
	return nil, nil
}

func (_this *PolicyRepository) ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error) {
	return nil, nil
}

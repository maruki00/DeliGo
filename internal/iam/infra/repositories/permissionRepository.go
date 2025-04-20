package repositories

import (
	"context"
	"deligo/internal/iam/infra/models"
)

type PermissionRepository struct{}

func (_this *PermissionRepository) FindByID(ctx context.Context, id string) (*models.Permission, error) {
	return nil, nil
}
func (_this *PermissionRepository) FindByPolicyID(ctx context.Context, policyID string) ([]*models.Permission, error) {
	return nil, nil
}
func (_this *PermissionRepository) Save(ctx context.Context, permission *models.Permission) error {
	return nil
}
func (_this *PermissionRepository) Delete(ctx context.Context, id string) error {
	return nil
}

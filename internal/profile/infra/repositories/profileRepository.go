package repositories

import (
	"context"
	"deligo/internal/profile/infra/models"
	pkgPostgres "deligo/pkg/postgres"
	"sync"
)

type ProfileRepository struct {
	sync.RWMutex
	db pkgPostgres.DBHandler
}

func NewProfileRepository(db pkgPostgres.DBHandler) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (_this *ProfileRepository) FindByID(ctx context.Context, id string) (*models.Policy, error) {
	return nil, nil
}
func (_this *ProfileRepository) FindByName(ctx context.Context, name string) (*models.Policy, error) {
	return nil, nil
}
func (_this *ProfileRepository) Save(ctx context.Context, policy *models.Policy) error {
	return nil
}
func (_this *ProfileRepository) Delete(ctx context.Context, id string) error {
	return nil
}
func (_this *ProfileRepository) ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error) {
	return nil, nil
}

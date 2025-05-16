package repositories

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/infra/models"
	pkgPostgres "deligo/pkg/postgres"
	"sync"
)

type ProfileRepository struct {
	sync.RWMutex
	db pkgPostgres.DBHandler
}

func NewProfileRepository(db pkgPostgres.DBHandler) contracts.IPorofileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (_this *ProfileRepository) Save(context.Context, *models.Profile) error {

	sql := `INSERT INTO profiles VALUES (?, ?, ?, ?, ?`

	return nil
}
func (_this *ProfileRepository) Disable(context.Context, *models.Profile) error {
	return nil
}
func (_this *ProfileRepository) FindByUserID(context.Context, string) (*models.Profile, error) {
	return nil
}
func (_this *ProfileRepository) Update(context.Context, string, map[string]any) error {
	return nil
}
func (_this *ProfileRepository) UpdateAvatar(context.Context, string, string) error {
	return nil
}

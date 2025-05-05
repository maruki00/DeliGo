package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	pkgPostgres "deligo/pkg/postgres"
)

type RoleRepository struct {
	db *pkgPostgres.DBHandler
}

func NewRoleRepository(db *pkgPostgres.DBHandler) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (_this *RoleRepository) Create(ctx context.Context, role entities.RoleEntity) error {

	return nil
}

func (_this *RoleRepository) GetByID(ctx context.Context, id string) (entities.RoleEntity, error) {

	return nil, nil
}

func (_this *RoleRepository) GetByName(ctx context.Context, name string) (entities.RoleEntity, error) {

	return nil, nil
}

func (_this *RoleRepository) List(ctx context.Context) ([]entities.RoleEntity, error) {

	return nil, nil
}

func (_this *RoleRepository) Delete(ctx context.Context, id string) error {

	return nil
}

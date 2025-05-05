package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *pkgPostgres.PGHandler
}

func NewRoleRepository(db *pkgPostgres.PGHandler) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (_this *RoleRepository) Save(ctx context.Context, role entities.RoleEntity) error {
	// sql := insert into roles (id, name) values ($1, $2)
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO roles (id, name) VALUES ($1, $2)`
		if err := tx.Exec(sql, role.GetID(), role.GetName()).Error; err != nil {
			return err
		}
		return nil
	})
	return err
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

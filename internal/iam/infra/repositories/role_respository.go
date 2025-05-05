package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
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

func (_this *RoleRepository) FindByID(ctx context.Context, id valueobjects.ID) (entities.RoleEntity, error) {
	var role models.Role
	if err := _this.db.DB.Where("id = ? ", string(id)).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (_this *RoleRepository) FindByName(ctx context.Context, name string) (entities.RoleEntity, error) {
	var role models.Role
	if err := _this.db.DB.Where("name = ? ", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (_this *RoleRepository) List(ctx context.Context, pagination shared_models.Pagination) ([]entities.RoleEntity, error) {

	var roles []*models.Role

	return nil, nil
}

func (_this *RoleRepository) Delete(ctx context.Context, id string) error {

	return nil
}

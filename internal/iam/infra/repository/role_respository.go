package repository

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/domain/entity"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	shared_model "github.com/maruki00/deligo/internal/shared/model"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

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

func (_this *RoleRepository) Save(ctx context.Context, role entity.RoleEntity) error {
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO roles (id, name, description) VALUES ($1, $2, $3)`
		if err := tx.Exec(sql, role.GetID(), role.GetName(), role.GetDescription()).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (_this *RoleRepository) FindByID(ctx context.Context, id valueobjects.ID) (entity.RoleEntity, error) {
	var role model.Role
	if err := _this.db.DB.Where("id = ? ", string(id)).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (_this *RoleRepository) FindByName(ctx context.Context, name string) (entity.RoleEntity, error) {
	var role model.Role
	if err := _this.db.DB.Where("name = ? ", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (_this *RoleRepository) List(ctx context.Context, pagination shared_model.Pagination) ([]*model.Role, error) {

	roles := make([]*model.Role, pagination.Limit)
	if err := _this.db.DB.Model(&model.Role{}).Limit(pagination.GetLimit()).Offset(pagination.GetOffset()).Find(&roles).Error; err != nil {
		return []*model.Role{}, err
	}
	return roles, nil
}

func (_this *PolicyRepository) AffectPolicy(ctx context.Context, id, role_id, policy_id string) error {

	return _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "roles_policies" ("id", "role_id", "policy_id") VALUES (?, ?, ?)`
		if err := tx.Exec(sql,
			id,
			role_id,
			policy_id,
		).Error; err != nil {
			return err
		}

		return nil
	})
}

func (_this *RoleRepository) Delete(ctx context.Context, id valueobjects.ID) error {

	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&model.Role{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

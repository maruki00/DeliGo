package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	"deligo/internal/iam/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *pkgPostgres.PGHandler
}

func NewPermissionRepository(db *pkgPostgres.PGHandler) *PermissionRepository {
	return &PermissionRepository{
		db: db,
	}
}

func (_this *PermissionRepository) Save(ctx context.Context, permission entities.PermissionEntity) error {
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "permissions" 
			("id", "name", "action", "description")
			VALUES (?, ?, ?, ?)`

		if err := tx.Exec(sql, permission.GetID(), permission.GetName(), permission.GetAction(), permission.GetDescription()).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (_this *PermissionRepository) FindByID(ctx context.Context, id string) (*models.Permission, error) {
	var permession models.Permission
	if err := _this.db.DB.Where("id = ? ", id).First(&permession).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
func (_this *PermissionRepository) FindByPolicyID(ctx context.Context, policyID string) ([]*models.Permission, error) {
	return nil, nil
}

func (_this *PermissionRepository) Delete(ctx context.Context, id string) error {
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? ", id).Delete(&models.Permission{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

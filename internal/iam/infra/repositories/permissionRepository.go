package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
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

func (_this *PermissionRepository) FindByPolicyID(ctx context.Context, policyID string, pagination shared_models.Pagination) ([]*models.Permission, error) {
	res, err := _this.db.DB.Raw(`
			SELECT * from permissions pr 
			LEFT JOIN policies_permissions pp ON pr.id = pp.permission_id 
			LEFT JOIN policies pl on pp.policies_id = pl.id
			where pr.id = ?
	`, policyID).Rows()
	if err != nil {
		return nil, err
	}

	for res.Next() {

	}

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

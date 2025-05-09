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
			LIMIT ? OFFSET ?
	`, policyID, pagination.GetLimit(), pagination.GetOffset()).Rows()
	if err != nil {
		return nil, err
	}
	permissions := make([]*models.Permission, pagination.GetLimit())
	var (
		_ID          string
		_Name        string
		_Action      string
		_Description string
	)
	index := 0
	for res.Next() {
		res.Scan(&_ID, &_Name, &_Action, &_Description)
		permissions[index] = &models.Permission{
			ID:          valueobjects.ID(_ID),
			Name:        _Name,
			Action:      _Action,
			Description: _Description,
		}
		index++
	}

	return permissions[:index], nil
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

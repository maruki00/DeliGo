package repository

import (
	"context"
	"github.com/maruki00/deligo/internal/iam/domain/entity"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

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

func (_this *PermissionRepository) Save(ctx context.Context, permission entity.PermissionEntity) error {
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

func (_this *PermissionRepository) FindByID(ctx context.Context, id string) (*model.Permission, error) {
	var permession model.Permission
	if err := _this.db.DB.Where("id = ? ", id).First(&permession).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (_this *PermissionRepository) FindByPolicyID(ctx context.Context, policyID string, pagination model.Pagination) ([]*model.Permission, error) {
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
	permissions := make([]*model.Permission, pagination.GetLimit())
	var (
		_ID          string
		_Name        string
		_Action      string
		_Description string
	)
	index := 0
	for res.Next() {
		res.Scan(&_ID, &_Name, &_Action, &_Description)
		permissions[index] = &model.Permission{
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
		if err := tx.Where("id = ? ", id).Delete(&model.Permission{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/domain/entity"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	"github.com/maruki00/deligo/pkg/pagination"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *pkgPostgres.PGHandler
}

func NewUserRepository(db *pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(ctx context.Context, entity entity.UserEntity) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "users" 
			("id", "username", "email", "tenant_id", "password", "password_changed_at", "is_active", "last_login", "mfa_enabled", "mfa_secret", "deleted_at")
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		if err := tx.Exec(sql,
			entity.GetID(),
			entity.GetUsername(),
			entity.GetEmail(),
			entity.GetTenantID(),
			entity.GetPassword(),
			entity.GetPasswordChangedAt(),
			entity.GetIsActive(),
			entity.GetLastLogin(),
			entity.GetMFAEnabled(),
			entity.GetMFASecret(),
			entity.GetDeletedAt(),
		).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id sharedvo.ID) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, id sharedvo.ID, fields map[string]interface{}) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.User{}).Where("id = ?", string(id)).Updates(fields).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id sharedvo.ID) (*model.User, error) {
	var user model.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) ListByTenant(ctx context.Context, tenantID sharedvo.ID, pagination pagination.Pagination) ([]*model.User, error) {
	users := make([]*model.User, pagination.Limit)
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.User{}).Where("tenant_id = ?", tenantID).Limit(pagination.GetLimit()).Offset(pagination.GetOffset()).Find(&users).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) AffectRole(ctx context.Context, id sharedvo.ID, role_id string) error {
	return ur.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "roles_policies" ("id", "role_id", "policy_id") VALUES (?, ?, ?)`
		if err := tx.Exec(sql,
			id,
			role_id,
			"i",
		).Error; err != nil {
			return err
		}
		return nil
	})
}

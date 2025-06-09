package repository

import (
	"context"
	"deligo/internal/iam/domain/entity"
	"deligo/internal/iam/infra/model"
	sharedvo "deligo/internal/shared/valueobject"
	"deligo/pkg/pagination"
	pkgPostgres "deligo/pkg/postgres"

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

// type IUserRepository interface {
// 	Save(context.Context, entity.UserEntity) error
// 	Delete(context.Context, sharedvo.ID) error
// 	Update(context.Context, sharedvo.ID, map[string]interface{}) error
// 	FindByID(context.Context, sharedvo.ID) (*model.User, error)
// 	FindByEmail(context.Context, string) (*model.User, error)
// 	FindByUsername(context.Context, string) (*model.User, error)
// 	ListByTenant(context.Context, sharedvo.ID, pagination.Pagination) ([]*model.User, error)
// 	AffectRole(context.Context, string) error
// }

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

func (_this *PolicyRepository) AffectRole(ctx context.Context, id, role_id string) error {

	return _this.db.DB.Transaction(func(tx *gorm.DB) error {
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

package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	shared_models "deligo/internal/shared/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"github.com/volcengine/volc-sdk-golang/service/codePipeline/models"
	"gorm.io/gorm"
)

/*
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error

	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	ListByTenant(ctx context.Context, tenantID string, pagination Pagination) ([]*models.User, error)

*/

type UserRepository struct {
	db pkgPostgres.PGHandler
}

func NewUserRepository(db pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(ctx context.Context, entity entities.UserEntity) error {

	//sql := `INSERT INTO users(id, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
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

func (ur *UserRepository) Delete(ctx context.Context, id string) error {
	//sql := `UPDATE users SET deleted_at = now(), updated_at = now() WHERE id = $1 `
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.User{}, map[string]interface{}{
			"id": id,
		}).Error; err != nil {
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

func (ur *UserRepository) Update(ctx context.Context, entity entities.UserEntity) error {
	// sql := `
	// 		UPDATE users
	// 		SET email=$1, role=$2, updated_at = now()
	// 		WHERE id=$3
	// 		AND deleted_at = NULL
	// 	`
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity).UpdateColumns(&entity).Error; err != nil {
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

func (ur *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		WHERE id = $1
	// 		LIMIT 1
	// 	`
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
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

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		WHERE id = $1
	// 		LIMIT 1
	// 	`
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user).Where("email", email).Error; err != nil {
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

func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		WHERE id = $1
	// 		LIMIT 1
	// 	`
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user).Where("username", username).Error; err != nil {
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

func (ur *UserRepository) ListByTenant(ctx context.Context, tenantID string, pagination shared_models.Pagination) ([]*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		-- WHERE deleted_at = NULL
	// 		OFFSET $1
	// 		LIMIT $2
	// 	`

	users := make([]*models.User, pagination.Limit)
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&users).Where("tenant_id", tenantID).Limit(pagination.GetLimit()).Offset(pagination.GetOffset()).Error; err != nil {
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

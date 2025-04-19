package repositories

import (
	"context"
	"deligo/internal/iam/domain/contracts"
	"deligo/internal/iam/domain/entities"
	"deligo/internal/iam/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type UserRepository struct {
	db pkgPostgres.PGHandler
}

func NewUserRepository(db pkgPostgres.PGHandler) contracts.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error) {

	//sql := `INSERT INTO users(id, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) (bool, error) {
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
		return false, err
	}

	return true, nil
}

func (ur *UserRepository) Update(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error) {
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
		return nil, err
	}

	return entity, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, userId string) (entities.UserEntity, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		WHERE id = $1
	// 		LIMIT 1
	// 	`
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userId).Error; err != nil {
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

func (ur *UserRepository) GetMany(ctx context.Context, offset, limit int32) ([]*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		-- WHERE deleted_at = NULL
	// 		OFFSET $1
	// 		LIMIT $2
	// 	`

	users := make([]*models.User, limit)
	offset = (offset - 1) * offset
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&users).Limit(int(limit)).Offset(int(offset)).Error; err != nil {
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

func (ur *UserRepository) Search(ctx context.Context, query string, offset, limit int32) ([]*models.User, error) {
	// sql := `
	// 		SELECT id,email,role
	// 		FROM users
	// 		WHERE (email = $1 OR role = $1)
	// 		OFFSET $2
	// 		LIMIT $3
	// 	`
	users := make([]*models.User, limit)
	offset = (offset - 1) * offset
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select("id", "email", "role").
			Where("email = ? OR role = ?", query).
			Offset(int(offset)).
			Limit(int(limit)).
			Find(&users).Error; err != nil {
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

package postgres

import (
	"context"
	"errors"

	"github.com/maruki00/deligo/internal/iam/domain/entity"
	"github.com/maruki00/deligo/internal/iam/infra/mapper"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *pkgPostgres.PGHandler
}

func NewUserRepo(db *pkgPostgres.PGHandler) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (_this *UserRepo) Save(ctx context.Context, entity *entity.User) error {
	tx := _this.db.DB.Save(entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (_this *UserRepo) FindByID(ctx context.Context, id string) (*entity.User, error) {
	user := entity.User{}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (_this *UserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user).Where("username = ?", username).Find(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (_this *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := entity.User{}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user).Where("email = ?", email).Find(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (_this *UserRepo) Delete(ctx context.Context, id string) error {
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.User{}, map[string]interface{}{
			"id": id,
		}).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (_this *UserRepo) StatusUpdate(ctx context.Context, id, status string) error {
	if status != "ban" && status != "active" {
		return errors.New("status not allowed")
	}
	if id != "1" {
		return errors.New("user not found")
	}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id", id).Update("status", status).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (_this *UserRepo) FindAll(ctx context.Context, page, offset int) ([]*entity.User, error) {
	var users []*entity.User
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.User{}).Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (_this *UserRepo) Update(ctx context.Context, entity *entity.User) error {
	user := mapper.ToGormModel(entity)
	tx := _this.db.DB.Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

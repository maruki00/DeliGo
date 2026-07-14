package postgres

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/infra/model"
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

func (_this *UserRepo) Save(ctx context.Context, entity *model.User) error {
	tx := _this.db.DB.Save(entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (_this *UserRepo) FindByID(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
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

func (_this *UserRepo) FindByUserReponame(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}
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

func (_this *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
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
		if err := tx.Delete(&model.User{}, map[string]interface{}{
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

func (_this *UserRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.User{}).Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (_this *UserRepo) FindByTenantID(ctx context.Context, tenantID string) ([]*model.User, error) {
	var users []*model.User
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.User{}).Where("tenant_id = ?", tenantID).Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (_this *UserRepo) FindByTenantIDAndRole(ctx context.Context, tenantID string, role string) ([]*model.User, error) {
	var users []*model.User
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.User{}).Where("tenant_id = ? AND role = ?", tenantID, role).Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (_this *UserRepo) Update(ctx context.Context, entity *model.User) error {
	tx := _this.db.DB.Save(entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

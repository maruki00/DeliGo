package repository

import (
	"context"
	"github.com/maruki00/deligo/internal/profile/domain/entity"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	db pkgPostgres.DBHandler
}

func NewProfileRepository(db pkgPostgres.DBHandler) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (_this *ProfileRepository) Save(ctx context.Context, entity entity.ProfileEntity) error {
	err := _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO profiles (id, user_id, full_name, avatar, bio) VALUES (?, ?, ?, ?, ?)`
		if err := tx.Exec(sql,
			entity.GetID(),
			entity.GetUserID(),
			entity.GetFullName(),
			entity.GetAvatar(),
			entity.GetBio(),
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

func (_this *ProfileRepository) Disable(ctx context.Context, id string) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Profile{}).Where("id = ?", id).UpdateColumn("is_active", false).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (_this *ProfileRepository) FindByUserID(ctx context.Context, id string) (*model.Profile, error) {
	var profile model.Profile
	err := _this.db.GetDB().Where("user_id = ?", id).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (_this *ProfileRepository) FindByID(ctx context.Context, id string) (*model.Profile, error) {
	var profile model.Profile
	err := _this.db.GetDB().Where("id = ?", id).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (_this *ProfileRepository) UpdateAvatar(ctx context.Context, id string, avatar string) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Profile{}).Where("id = ?", id).UpdateColumn("avatar", avatar).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (_this *ProfileRepository) Update(ctx context.Context, id string, fields map[string]any) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Profile{}).Where("id = ?", id).UpdateColumns(fields).Error; err != nil {
			return err
		}
		return nil
	})
}

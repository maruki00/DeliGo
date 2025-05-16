package repositories

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/domain/entities"
	"deligo/internal/profile/infra/models"
	"deligo/internal/prrofile/domain/entities"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	db pkgPostgres.DBHandler
}

func NewProfileRepository(db pkgPostgres.DBHandler) contracts.IPorofileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (_this *ProfileRepository) Save(context.Context, entity *entities.ProfileEntity) error {



	r := _this.db.DB.Transaction(func(tx *gorm.DB) error {

		sql := `INSERT INTO profiles (id, user_name, full_name, avatar, bio) VALUES (?, ?, ?, ?, ?)`

		if err := tx.Exec(sql,
			entity.GetI(),
			entity.GetUserI(),
			entity.GetFullNam(),
			entity.GetAvata(),
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

	return nil
}
func (_this *ProfileRepository) Disable(context.Context, *models.Profile) error {
	return nil
}
func (_this *ProfileRepository) FindByUserID(context.Context, string) (*models.Profile, error) {
	return nil
}
func (_this *ProfileRepository) Update(context.Context, string, map[string]any) error {
	return nil
}
func (_this *ProfileRepository) UpdateAvatar(context.Context, string, string) error {
	return nil
}

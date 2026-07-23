package repository

import (
	"context"
	"errors"

	"github.com/maruki00/deligo/internal/user/domain"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Check(ctx context.Context, token string) error {
	var auth domain.Auth
	if err := r.db.WithContext(ctx).First(&auth, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrUserNotFound
		}

		return err
	}
	if auth.Token == "" {
		return errors.New("record not found.")
	}
	return nil
}

func (r *AuthRepository) CreateToken(ctx context.Context, email, token string) error {
	if err := r.db.Model(&domain.Auth{}).WithContext(ctx).Where("email", email).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.db.Model(&domain.Auth{}).WithContext(ctx).Save(&domain.Auth{
		Token: token,
		Email: email,
	}).Error; err != nil {
		return err
	}
	return nil
}

package mapper

import (
	"github.com/maruki00/deligo/internal/iam/domain/entity"
	"github.com/maruki00/deligo/internal/iam/infra/model"
)

func toEntity(m *model.UserGormModel) *entity.User {
	return &entity.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Phone:        m.Phone,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		Role:         entity.UserRole(m.Role),
		Status:       entity.UserStatus(m.Status),
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func ToGormModel(u *entity.User) *model.UserGormModel {
	return &model.UserGormModel{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Phone:        u.Phone,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Role:         string(u.Role),
		Status:       string(u.Status),
	}
}

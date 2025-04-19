package entities

import (
	"deligo/internal/iam/infra/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity interface {
	SetID(ID uuid.UUID)
	SetUsername(Username string)
	SetEmail(Email string)
	SetTenantID(TenantID string)
	SetPassword(Password string)
	SetPasswordHash(PasswordHash string)
	SetPasswordChangedAt(PasswordChangedAt *time.Time)
	SetIsActive(IsActive bool)
	SetLastLogin(LastLogin *time.Time)
	SetMFAEnabled(MFAEnabled bool)
	SetMFASecret(MFASecret string)
	SetDeletedAt(DeletedAt gorm.DeletedAt)
	SetCreatedAt(CreatedAt time.Time)
	SetUpdatedAt(UpdatedAt time.Time)
	SetProfile(Profile models.Profile)
	SetGroups(Groups []models.Group)
	//hello world
	GetID() uuid.UUID
	GetUsername() string
	GetEmail() string
	GetTenantID() string
	GetPassword() string
	GetPasswordHash() string
	GetPasswordChangedAt() *time.Time
	GetIsActive() bool
	GetLastLogin() *time.Time
	GetMFAEnabled() bool
	GetMFASecret() string
	GetDeletedAt() gorm.DeletedAt
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetProfile() models.Profile
	GetGroups() []models.Group
}

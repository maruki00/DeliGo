package entities

import (
	"deligo/internal/iam/infra/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity interface {
	// SetID(ID string)
	// SetEmail(Email string)
	// SetPassword(Password string)
	// SetRole(Role string)
	// SetDeletedAt(DeletedAt *time.Time)
	// SetCreatedAt(CreatedAt *time.Time)
	// SetUpdatedAt(UpdatedAt *time.Time)
	// GetID() string
	// GetEmail() string
	// GetPassword() string
	// GetRole() string
	// GetDeletedAt() *time.Time
	// GetCreatedAt() *time.Time
	// GetUpdatedAt() *time.Time
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
	SetRoles(Roles []models.Role)
	SetDeletedAt(DeletedAt gorm.DeletedAt)
	SetCreatedAt(CreatedAt time.Time)
	SetUpdatedAt(UpdatedAt time.Time)
	SetProfile(Profile models.Profile)
	SetGroups(Groups []models.Group)
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
	GetRoles() []models.Role
	GetDeletedAt() gorm.DeletedAt
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetProfile() models.Profile
	GetGroups() []models.Group
}

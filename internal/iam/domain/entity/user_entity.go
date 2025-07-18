package entity

import (
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	"time"

	"gorm.io/gorm"
)

type UserEntity interface {
	SetID(ID valueobjects.ID)
	SetUsername(Username string)
	SetEmail(Email string)
	SetTenantID(TenantID valueobjects.ID)
	SetPassword(Password valueobjects.Password)
	SetPasswordChangedAt(PasswordChangedAt *time.Time)
	SetIsActive(IsActive bool)
	SetLastLogin(LastLogin *time.Time)
	SetMFAEnabled(MFAEnabled bool)
	SetMFASecret(MFASecret string)
	SetDeletedAt(DeletedAt gorm.DeletedAt)
	SetCreatedAt(CreatedAt time.Time)
	SetUpdatedAt(UpdatedAt time.Time)

	GetID() valueobjects.ID
	GetUsername() string
	GetEmail() string
	GetTenantID() valueobjects.ID
	GetPassword() valueobjects.Password
	GetPasswordChangedAt() *time.Time
	GetIsActive() bool
	GetLastLogin() *time.Time
	GetMFAEnabled() bool
	GetMFASecret() string
	GetDeletedAt() gorm.DeletedAt
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetPolicies() []*model.Policy
}

package models

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                valueobjects.ID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username          string                `json:"user_name" gorm:"type:varchar(255)"`
	Email             string                `json:"email" gorm:"type:varchar(255)"`
	TenantID          valueobjects.ID       `json:"tenant_id" gorm:"type:varchar(255)"`
	Password          valueobjects.Password `json:"password" gorm:"type:varchar(255)"`
	PasswordChangedAt *time.Time            `json:"password_changed_at" gorm:"type:varchar(255)"`
	IsActive          bool                  `json:"is_active" gorm:"type:int;default:0"`
	LastLogin         *time.Time            `json:"last_login" `
	MFAEnabled        bool                  `json:"mfa_enabled" gorm:"default:0"`
	MFASecret         string                `json:"mfa_secret"`
	Profile           *Profile              `json:"profile" gorm:"foreignKey:UserID"`
	Groups            []*Group              `json:"groups" gorm:"many2many:user_groups;"`
	Policies          []*Policy             `json:"policies" gorm:"many2many:user_policies;"`
	DeletedAt         gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
	CreatedAt         time.Time             `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt         time.Time             `json:"updated_at" gorm:"not null;default:now()"`
}

func (_this *User) SetID(ID valueobjects.ID) {
	_this.ID = ID
}
func (_this *User) SetUsername(Username string) {
	_this.Username = Username
}
func (_this *User) SetEmail(Email string) {
	_this.Email = Email
}
func (_this *User) SetTenantID(TenantID valueobjects.ID) {
	_this.TenantID = TenantID
}
func (_this *User) SetPassword(Password valueobjects.Password) {
	_this.Password = Password
}
func (_this *User) SetPasswordChangedAt(PasswordChangedAt *time.Time) {
	_this.PasswordChangedAt = PasswordChangedAt
}
func (_this *User) SetIsActive(IsActive bool) {
	_this.IsActive = IsActive
}
func (_this *User) SetLastLogin(LastLogin *time.Time) {
	_this.LastLogin = LastLogin
}
func (_this *User) SetMFAEnabled(MFAEnabled bool) {
	_this.MFAEnabled = MFAEnabled
}
func (_this *User) SetMFASecret(MFASecret string) {
	_this.MFASecret = MFASecret
}
func (_this *User) SetDeletedAt(DeletedAt gorm.DeletedAt) {
	_this.DeletedAt = DeletedAt
}
func (_this *User) SetCreatedAt(CreatedAt time.Time) {
	_this.CreatedAt = CreatedAt
}
func (_this *User) SetUpdatedAt(UpdatedAt time.Time) {
	_this.UpdatedAt = UpdatedAt
}

func (_this *User) GetID() valueobjects.ID {
	return _this.ID
}
func (_this *User) GetUsername() string {
	return _this.Username
}
func (_this *User) GetEmail() string {
	return _this.Email
}
func (_this *User) GetTenantID() valueobjects.ID {
	return _this.TenantID
}
func (_this *User) GetPassword() valueobjects.Password {
	return _this.Password
}
func (_this *User) GetPasswordChangedAt() *time.Time {
	return _this.PasswordChangedAt
}
func (_this *User) GetIsActive() bool {
	return _this.IsActive
}
func (_this *User) GetLastLogin() *time.Time {
	return _this.LastLogin
}
func (_this *User) GetMFAEnabled() bool {
	return _this.MFAEnabled
}
func (_this *User) GetMFASecret() string {
	return _this.MFASecret
}
func (_this *User) GetDeletedAt() gorm.DeletedAt {
	return _this.DeletedAt
}
func (_this *User) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *User) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

//	func (_this *User) GetProfile() *Profile {
//		return _this.Profile
//	}
func (_this *User) GetGroups() []*Group {
	return _this.Groups
}
func (_this *User) GetPolicies() []*Policy {
	return _this.Policies
}

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID
	Username          string
	Email             string
	TenantID          string
	Password          string
	PasswordHash      string
	PasswordChangedAt *time.Time
	IsActive          bool
	LastLogin         *time.Time
	MFAEnabled        bool
	MFASecret         string
	Roles             []Role
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Profile           Profile `gorm:"foreignKey:UserID"`
	Groups            []Group `gorm:"many2many:user_groups;"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.ID == uuid.Nil {
// 		u.ID = uuid.New()
// 	}
// 	return nil
// }

func (_this *User) SetID(ID uuid.UUID) {
	_this.ID = ID
}
func (_this *User) SetUsername(Username string) {
	_this.Username = Username
}
func (_this *User) SetEmail(Email string) {
	_this.Email = Email
}
func (_this *User) SetTenantID(TenantID string) {
	_this.TenantID = TenantID
}
func (_this *User) SetPassword(Password string) {
	_this.Password = Password
}
func (_this *User) SetPasswordHash(PasswordHash string) {
	_this.PasswordHash = PasswordHash
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
func (_this *User) SetRoles(Roles []Role) {
	_this.Roles = Roles
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
func (_this *User) SetProfile(Profile Profile) {
	_this.Profile = Profile
}
func (_this *User) SetGroups(Groups []Group) {
	_this.Groups = Groups
}

func (_this *User) GetID() uuid.UUID {
	return _this.ID
}
func (_this *User) GetUsername() string {
	return _this.Username
}
func (_this *User) GetEmail() string {
	return _this.Email
}
func (_this *User) GetTenantID() string {
	return _this.TenantID
}
func (_this *User) GetPassword() string {
	return _this.Password
}
func (_this *User) GetPasswordHash() string {
	return _this.PasswordHash
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
func (_this *User) GetRoles() []Role {
	return _this.Roles
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
func (_this *User) GetProfile() Profile {
	return _this.Profile
}
func (_this *User) GetGroups() []Group {
	return _this.Groups
}

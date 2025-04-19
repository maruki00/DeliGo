package models

// import (
// 	"time"
// )

// type User struct {
// 	ID        string
// 	Email     string
// 	Password  string
// 	Role      string
// 	DeletedAt *time.Time
// 	CreatedAt *time.Time
// 	UpdatedAt *time.Time
// }

// func (_this *User) SetID(ID string) {
// 	_this.ID = ID
// }

// func (_this *User) SetEmail(Email string) {
// 	_this.Email = Email
// }

// func (_this *User) SetPassword(Password string) {
// 	_this.Password = Password
// }

// func (_this *User) SetRole(Role string) {
// 	_this.Role = Role
// }

// func (_this *User) SetDeletedAt(DeletedAt *time.Time) {
// 	_this.DeletedAt = DeletedAt
// }

// func (_this *User) SetCreatedAt(CreatedAt *time.Time) {
// 	_this.CreatedAt = CreatedAt
// }

// func (_this *User) SetUpdatedAt(UpdatedAt *time.Time) {
// 	_this.UpdatedAt = UpdatedAt
// }

// func (_this *User) GetID() string {
// 	return _this.ID
// }

// func (_this *User) GetEmail() string {
// 	return _this.Email
// }

// func (_this *User) GetPassword() string {
// 	return _this.Password
// }

// func (_this *User) GetRole() string {
// 	return _this.Role
// }

// func (_this *User) GetDeletedAt() *time.Time {
// 	return _this.DeletedAt
// }

// func (_this *User) GetCreatedAt() *time.Time {
// 	return _this.CreatedAt
// }

// func (_this *User) GetUpdatedAt() *time.Time {
// 	return _this.UpdatedAt
// }

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username          string         `gorm:"type:varchar(32);not null;uniqueIndex:idx_username_tenant"`
	Email             string         `gorm:"type:varchar(255);not null;uniqueIndex:idx_email_tenant"`
	TenantID          string         `gorm:"type:varchar(32);not null;index"`
	Password          string         `gorm:"type:varchar(255);not null"`
	PasswordChangedAt *time.Time     `gorm:"default:null"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	CreatedAt         time.Time      `gorm:"not null;default:now()"`
	UpdatedAt         time.Time      `gorm:"not null;default:now()"`

	// Associations
	Profile Profile `gorm:"foreignKey:UserID"`
	Groups  []Group `gorm:"many2many:user_groups;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

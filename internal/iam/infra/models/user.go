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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroup struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    string         `gorm:"type:varchar(32);not null;index:idx_user_group"`
	GroupID   string         `gorm:"type:varchar(32);not null;index:idx_user_group"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"not null;default:now()"`
	UpdatedAt time.Time      `gorm:"not null;default:now()"`
}

func (ug *UserGroup) BeforeCreate(tx *gorm.DB) error {
	if ug.ID == uuid.Nil {
		ug.ID = uuid.New()
	}
	return nil
}

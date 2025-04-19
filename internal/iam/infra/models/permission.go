package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text;not null"`
	PolicyID    string         `gorm:"type:varchar(255);not null;index"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedAt   time.Time      `gorm:"not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"not null;default:now()"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

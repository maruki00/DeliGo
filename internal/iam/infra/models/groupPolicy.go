package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupPolicy struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	GroupID   string         `gorm:"type:varchar(32);not null;index:idx_group_policy"`
	PolicyID  string         `gorm:"type:varchar(32);not null;index:idx_group_policy"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"not null;default:now()"`
	UpdatedAt time.Time      `gorm:"not null;default:now()"`
}

func (gp *GroupPolicy) BeforeCreate(tx *gorm.DB) error {
	if gp.ID == uuid.Nil {
		gp.ID = uuid.New()
	}
	return nil
}

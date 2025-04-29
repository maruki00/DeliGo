package models

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
	"time"

	"gorm.io/gorm"
)

type Policy struct {
	ID          valueobjects.ID `gorm:"type:uuid;primaryKey"`
	Name        string          `gorm:"type:varchar(255);not null"`
	GroupID     string          `gorm:"type:varchar(32);not null;index"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	CreatedAt   time.Time       `gorm:"not null;default:now()"`
	UpdatedAt   time.Time       `gorm:"not null;default:now()"`
	Permissions []*Permission   `gorm:"foreignKey:PolicyID"`
	Groups      []*Group        `gorm:"many2many:group_policies;"`
}

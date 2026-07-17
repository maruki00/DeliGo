package model

import (
	"time"
)

type UserGormModel struct {
	ID           string    `gorm:"primaryKey;type:varchar(36)"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Phone        string    `gorm:"uniqueIndex;type:varchar(50);not null"`
	FirstName    string    `gorm:"type:varchar(100)"`
	LastName     string    `gorm:"type:varchar(100)"`
	Role         string    `gorm:"type:varchar(50);not null"`
	Status       string    `gorm:"type:varchar(50);default:'active'"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserGormModel) TableName() string {
	return "users"
}

type CasbinRuleGormModel struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"type:varchar(100);not null;index:idx_casbin_rule,priority:1"`
	V0    string `gorm:"type:varchar(100);index:idx_casbin_rule,priority:2"`
	V1    string `gorm:"type:varchar(100);index:idx_casbin_rule,priority:3"`
	V2    string `gorm:"type:varchar(100);index:idx_casbin_rule,priority:4"`
	V3    string `gorm:"type:varchar(100)"`
	V4    string `gorm:"type:varchar(100)"`
	V5    string `gorm:"type:varchar(100)"`
}

func (CasbinRuleGormModel) TableName() string {
	return "casbin_rule"
}

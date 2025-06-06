package model

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
}

func (_this *Policy) SetID(_ID valueobjects.ID) {
	_this.ID = _ID
}
func (_this *Policy) SetName(_Name string) {
	_this.Name = _Name
}
func (_this *Policy) GetID() valueobjects.ID {
	return _this.ID
}
func (_this *Policy) GetName() string {
	return _this.Name
}
func (_this *Policy) GetGroupID() string {
	return _this.GroupID
}
func (_this *Policy) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *Policy) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}
func (_this *Policy) GetPermissions() []*Permission {
	return _this.Permissions
}

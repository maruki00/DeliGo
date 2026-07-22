package models

import (
	"time"
)

type Restaurant struct {
	ID        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	OwnerID   string    `gorm:"type:varchar(36);not null;index:idx_owner" json:"owner_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	IsOpen    bool      `gorm:"type:boolean;default:false" json:"is_open"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Products  []Product `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE" json:"products,omitempty"`
}

type Product struct {
	ID           string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	RestaurantID string    `gorm:"type:varchar(36);not null" json:"restaurant_id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	Price        float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	IsAvailable  bool      `gorm:"type:boolean;default:true" json:"is_available"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

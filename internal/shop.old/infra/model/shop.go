package model

import (
	"errors"
	"time"
)

var (
	ErrShopNameRequired = errors.New("shop name is required")
	ErrShopNotFound     = errors.New("shop not found")
)

type Shop struct {
	ID        string `gorm:"primaryKey;type:uuid;not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	Address   string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(50)"`
	OwnerID   string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewShop(id, name, address, phone, ownerID string) (*Shop, error) {
	if name == "" {
		return nil, ErrShopNameRequired
	}
	now := time.Now().UTC()
	return &Shop{
		ID:        id,
		Name:      name,
		Address:   address,
		Phone:     phone,
		OwnerID:   ownerID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (s *Shop) UpdateDetails(name, address, phone string) error {
	if name == "" {
		return ErrShopNameRequired
	}
	s.Name = name
	s.Address = address
	s.Phone = phone
	s.UpdatedAt = time.Now().UTC()
	return nil
}

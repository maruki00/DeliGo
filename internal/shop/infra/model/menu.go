package model

import (
	"errors"
	"time"
)

var (
	ErrMenuNameRequired  = errors.New("menu name is required")
	ErrMenuNotFound      = errors.New("menu not found")
	ErrItemAlreadyExists = errors.New("product already exists in menu")
	ErrItemNotFound      = errors.New("menu item not found")
)

type Item struct {
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
}

type Menu struct {
	ID        string `gorm:"primaryKey;type:uuid;not null"`
	ShopID    string `gorm:"type:varchar(255);not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	Items     []Item `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMenu(id, shopID, name string) (*Menu, error) {
	if name == "" {
		return nil, ErrMenuNameRequired
	}
	now := time.Now().UTC()
	return &Menu{
		ID:        id,
		ShopID:    shopID,
		Name:      name,
		Items:     []Item{},
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (m *Menu) AddItem(productID string, price float64) error {
	for _, it := range m.Items {
		if it.ProductID == productID {
			return ErrItemAlreadyExists
		}
	}
	m.Items = append(m.Items, Item{ProductID: productID, Price: price})
	m.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Menu) RemoveItem(productID string) error {
	for i, it := range m.Items {
		if it.ProductID == productID {
			m.Items = append(m.Items[:i], m.Items[i+1:]...)
			m.UpdatedAt = time.Now().UTC()
			return nil
		}
	}
	return ErrItemNotFound
}

package model

import (
	"errors"
	"time"
)

var (
	ErrProductNameRequired = errors.New("product name is required")
	ErrInvalidPrice        = errors.New("price must be greater than zero")
	ErrInsufficientStock   = errors.New("insufficient stock")
	ErrProductNotFound     = errors.New("product not found")
)

type Product struct {
	ID          string
	ShopID      string
	Name        string
	Description string
	Price       float64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(id, shopID, name, description string, price float64, stock int) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameRequired
	}
	if price <= 0 {
		return nil, ErrInvalidPrice
	}
	now := time.Now().UTC()
	return &Product{
		ID:          id,
		ShopID:      shopID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (p *Product) UpdateDetails(name, description string, price float64) error {
	if name == "" {
		return ErrProductNameRequired
	}
	if price <= 0 {
		return ErrInvalidPrice
	}
	p.Name = name
	p.Description = description
	p.Price = price
	p.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Product) AdjustStock(delta int) error {
	if p.Stock+delta < 0 {
		return ErrInsufficientStock
	}
	p.Stock += delta
	p.UpdatedAt = time.Now().UTC()
	return nil
}

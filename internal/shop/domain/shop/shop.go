package shop

import "errors"

var (
	ErrShopNameRequired = errors.New("shop name is required")
)

type Shop struct {
	ID      string
	Name    string
	Address string
	Phone   string
	OwnerID string
}

func NewShop(id, name, address, phone, ownerID string) (*Shop, error) {
	if name == "" {
		return nil, ErrShopNameRequired
	}

	return &Shop{
		ID:      id,
		Name:    name,
		Address: address,
		Phone:   phone,
		OwnerID: ownerID,
	}, nil
}

func (s *Shop) UpdateDetails(name, address, phone string) error {
	if name == "" {
		return ErrShopNameRequired
	}

	s.Name = name
	s.Address = address
	s.Phone = phone
	return nil
}

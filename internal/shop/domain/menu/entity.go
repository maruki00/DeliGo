package menu

import "errors"

var (
	ErrMenuNameRequired  = errors.New("menu name is required")
	ErrItemAlreadyExists = errors.New("product already exists in menu")
	ErrItemNotFound      = errors.New("menu item not found")
)

type Item struct {
	ProductID string
	Price     float64
}

type Menu struct {
	ID     string
	ShopID string
	Name   string
	Items  []Item
}

func NewMenu(id, shopID, name string) (*Menu, error) {
	if name == "" {
		return nil, ErrMenuNameRequired
	}

	return &Menu{
		ID:     id,
		ShopID: shopID,
		Name:   name,
		Items:  []Item{},
	}, nil
}

func (m *Menu) AddItem(productID string, price float64) error {
	for _, item := range m.Items {
		if item.ProductID == productID {
			return ErrItemAlreadyExists
		}
	}

	m.Items = append(m.Items, Item{ProductID: productID, Price: price})
	return nil
}

func (m *Menu) RemoveItem(productID string) error {
	for i, item := range m.Items {
		if item.ProductID == productID {
			m.Items = append(m.Items[:i], m.Items[i+1:]...)
			return nil
		}
	}

	return ErrItemNotFound
}

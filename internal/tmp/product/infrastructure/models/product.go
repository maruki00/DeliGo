package models

type Product struct {
	Id    int     `json:"id"`
	Label string  `json:"label"`
	Price float32 `json:"price"`
	Type  string  `json:"type"`
}

func (p *Product) GetLabel() string {
	return p.Label
}
func (p *Product) GetId() int {
	return p.Id
}
func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) GetType() string {
	return p.Type
}

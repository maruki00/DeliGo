package valueobjects

import (
	"fmt"
)

type PriceVauleObject struct {
	price float32
}

func NewPriceValueObject(price float32) *PriceVauleObject {
	if price == 0.0 {
		panic("invalid price")
	}
	return &PriceVauleObject{
		price: price,
	}
}

func (p *PriceVauleObject) String() string {
	return fmt.Sprintf("%f", p.price)
}

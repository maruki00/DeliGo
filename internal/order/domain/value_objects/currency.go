package valueobjects

import (
	"fmt"
	"strings"
)

const (
	USD = "usd"
	MAD = "mad"
)

const (
	USD_MAD = 0.1
)

type Currency struct {
	Code    string
	Ammount float32
}

func (_this *Currency) Value() float32 {
	code := strings.ToLower(_this.Code)
	switch code {
	case USD:
		return _this.Ammount
	case MAD:
		return _this.Ammount * USD_MAD
	default:
		return 0
	}
}

func (_this *Currency) String() string {
	return fmt.Sprintf("%v", _this.Value())
}

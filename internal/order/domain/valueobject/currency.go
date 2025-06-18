package valueobject

import (
	"fmt"
	"strings"

	enums "github.com/maruki00/deligo/internal/order/domain/enum"
)

type Currency struct {
	Code    string
	Ammount float32
}

func (_this *Currency) Value() float32 {
	code := strings.ToLower(_this.Code)
	switch code {
	case enums.USD:
		return _this.Ammount
	case enums.MAD:
		return _this.Ammount * enums.USD_MAD
	default:
		return 0
	}
}

func (_this *Currency) String() string {
	return fmt.Sprintf("%v", _this.Value())
}

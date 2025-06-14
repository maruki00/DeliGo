package command

import (
	"time"

	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type UpdateShopCommand struct {
	ID       sharedvo.ID
	ShopName string
	OpenAt   time.Time
	CloseAt  time.Time
}

func (_this *UpdateShopCommand) Name() string {
	return "UpdateShopCommand"
}

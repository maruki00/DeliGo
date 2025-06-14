package command

import (
	sharedvo "github.com/maruki00/deligo/internal/shared/valueobject"
	"time"
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

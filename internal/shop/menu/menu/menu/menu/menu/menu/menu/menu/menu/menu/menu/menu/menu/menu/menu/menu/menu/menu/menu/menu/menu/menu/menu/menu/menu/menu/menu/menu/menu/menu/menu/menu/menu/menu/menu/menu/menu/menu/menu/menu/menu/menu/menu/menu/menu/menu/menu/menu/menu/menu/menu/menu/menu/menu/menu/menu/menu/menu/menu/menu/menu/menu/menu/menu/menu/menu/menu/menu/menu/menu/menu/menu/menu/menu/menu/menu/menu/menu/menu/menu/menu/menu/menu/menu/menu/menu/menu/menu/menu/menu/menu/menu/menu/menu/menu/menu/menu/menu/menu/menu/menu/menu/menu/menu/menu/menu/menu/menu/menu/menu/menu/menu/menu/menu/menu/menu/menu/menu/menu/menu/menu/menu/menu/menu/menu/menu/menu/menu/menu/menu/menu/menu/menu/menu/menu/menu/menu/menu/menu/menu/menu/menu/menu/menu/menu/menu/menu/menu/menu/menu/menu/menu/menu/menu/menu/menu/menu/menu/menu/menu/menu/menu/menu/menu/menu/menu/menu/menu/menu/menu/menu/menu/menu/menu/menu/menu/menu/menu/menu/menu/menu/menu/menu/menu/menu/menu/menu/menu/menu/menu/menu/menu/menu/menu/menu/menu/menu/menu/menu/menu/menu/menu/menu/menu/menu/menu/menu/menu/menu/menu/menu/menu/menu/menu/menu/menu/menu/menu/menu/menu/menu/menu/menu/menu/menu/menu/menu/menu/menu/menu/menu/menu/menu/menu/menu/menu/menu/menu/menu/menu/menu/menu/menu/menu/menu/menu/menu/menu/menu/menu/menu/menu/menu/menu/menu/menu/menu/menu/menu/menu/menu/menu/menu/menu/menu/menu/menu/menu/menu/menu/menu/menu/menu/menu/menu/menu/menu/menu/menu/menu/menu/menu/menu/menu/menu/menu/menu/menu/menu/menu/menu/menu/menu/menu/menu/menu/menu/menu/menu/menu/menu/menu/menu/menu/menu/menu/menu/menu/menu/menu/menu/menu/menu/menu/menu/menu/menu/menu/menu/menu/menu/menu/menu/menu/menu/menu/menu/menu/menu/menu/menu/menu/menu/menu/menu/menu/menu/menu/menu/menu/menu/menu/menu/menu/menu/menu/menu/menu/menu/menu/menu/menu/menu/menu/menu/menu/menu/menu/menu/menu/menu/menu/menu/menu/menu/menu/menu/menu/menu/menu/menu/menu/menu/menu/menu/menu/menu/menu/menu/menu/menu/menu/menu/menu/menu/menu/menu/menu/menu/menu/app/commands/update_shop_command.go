package commands

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"time"
)

type UpdateShopCommand struct {
	ID       shared_valueobject.ID
	ShopName string
	OpenAt   time.Time
	CloseAt  time.Time
}

func (_this *UpdateShopCommand) Name() string {
	return "UpdateShopCommand"
}

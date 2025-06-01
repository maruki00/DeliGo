package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type CreateShopCommand struct {
	ID       shared_valueobject.ID
	ShopName string
}

func (_this *CreateShopCommand) Name() string {
	return "CreateShopCommand"
}

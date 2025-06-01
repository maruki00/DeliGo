package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type CloseShopCommand struct {
	ID shared_valueobject.ID
}

func (_this *CloseShopCommand) Name() string {
	return "CloseShopCommand"
}

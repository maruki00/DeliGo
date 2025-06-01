package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type DeleteShopCommand struct {
	ID shared_valueobject.ID
}

func (_this *DeleteShopCommand) Name() string {
	return "DeleteShopCommand"
}

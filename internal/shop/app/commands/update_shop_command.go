package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type UpdateShopCommand struct {
	ID shared_valueobject.ID
}

func (_this *UpdateShopCommand) Name() string {
	return "UpdateShopCommand"
}

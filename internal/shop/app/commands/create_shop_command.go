package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type CreateShopCommand struct {
	ID    shared_valueobject.ID
	Label string
}

func (_this *CreateShopCommand) Name() string {
	return "CreateShopCommand"
}

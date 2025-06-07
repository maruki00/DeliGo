package command

import sharedvo "deligo/internal/shared/valueobject"

type CloseShopCommand struct {
	ID sharedvo.ID
}

func (_this *CloseShopCommand) Name() string {
	return "CloseShopCommand"
}

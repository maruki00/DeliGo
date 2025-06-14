package command

import sharedvo "github.com/maruki00/deligo/internal/shared/valueobject"

type CloseShopCommand struct {
	ID sharedvo.ID
}

func (_this *CloseShopCommand) Name() string {
	return "CloseShopCommand"
}

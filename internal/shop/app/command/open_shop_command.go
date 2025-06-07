package command

import sharedvo "deligo/internal/shared/valueobject"

type OpenShopCommand struct {
	ID sharedvo.ID
}

func (_this *OpenShopCommand) Name() string {
	return "OpenShopCommand"
}

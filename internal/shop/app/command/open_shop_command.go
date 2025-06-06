package commands

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type OpenShopCommand struct {
	ID shared_valueobject.ID
}

func (_this *OpenShopCommand) Name() string {
	return "OpenShopCommand"
}

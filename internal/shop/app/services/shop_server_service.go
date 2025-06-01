package services

import pkgCqrs "deligo/pkg/cqrs"

type ShopServerService struct {
	cmdBus *pkgCqrs.CommandBus
	qryBus *pkgCqrs.QueryBus
}

func NewShopServerService(cmdBus *pkgCqrs.CommandBus, qryBus *pkgCqrs.QueryBus) *ShopServerService {
	srv := ShopServerService{
		cmdBus: cmdBus,
		qryBus: qryBus,
	}
	return &srv
}

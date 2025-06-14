package service

import (
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type MenuServerService struct {
	cmdBus *pkgCqrs.CommandBus
	qryBus *pkgCqrs.QueryBus
}

func NewShopServerService(cmdBus *pkgCqrs.CommandBus, qryBus *pkgCqrs.QueryBus) *MenuServerService {
	srv := MenuServerService{
		cmdBus: cmdBus,
		qryBus: qryBus,
	}
	return &srv
}

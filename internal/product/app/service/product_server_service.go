package services

import (
	profile_grpc "github.com/maruki00/deligo/internal/profile/infra/grpc/profile"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type ProductServerService struct {
	profile_grpc.UnimplementedProfileServiceServer
	cmdBus *pkgCqrs.CommandBus
	qryBus *pkgCqrs.QueryBus
}

func NewProductService(cmdBus *pkgCqrs.CommandBus, qryBus *pkgCqrs.QueryBus) *ProductServerService {
	return &ProductServerService{
		qryBus: qryBus,
		cmdBus: cmdBus,
	}
}

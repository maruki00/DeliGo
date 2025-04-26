package groupHandlers

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetUserGroupsHandler struct {
	ID string
}

func (_this *GetUserGroupsHandler) Handle(ctx context.Context, command pkgCqrs.Query) (interface{}, error) {
	return nil, nil
}

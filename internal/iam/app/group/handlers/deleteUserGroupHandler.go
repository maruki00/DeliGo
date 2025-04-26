package groupHandlers

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type DeleteUserGroupCommand struct {
}

func (_this *DeleteUserGroupCommand) Handle(ctx context.Context, command pkgCqrs.Command) error {
	return nil
}

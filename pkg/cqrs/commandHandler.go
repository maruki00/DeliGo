package pkgCqrs

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, command Command) error
}

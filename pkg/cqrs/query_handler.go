package pkgCqrs

import "context"

type QueryHandler interface {
	Handle(ctx context.Context, query Query) (interface{}, error)
}

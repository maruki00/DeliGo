package pkgCqrs

import (
	"context"
	"fmt"
)

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (b *QueryBus) Register(query Query, handler QueryHandler) {
	queryName := query.Name()
	b.handlers[queryName] = handler
}

func (b *QueryBus) Dispatch(ctx context.Context, query Query) (interface{}, error) {
	queryName := query.Name()
	handler, exists := b.handlers[queryName]
	if !exists {
		return nil, fmt.Errorf("no handler registered for query %s", queryName)
	}
	return handler.Handle(ctx, query)
}

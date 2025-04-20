package pkgCqrs

import (
	"fmt"
	"reflect"
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
	queryName := reflect.TypeOf(query).Name()
	b.handlers[queryName] = handler
}

func (b *QueryBus) Dispatch(query Query) (interface{}, error) {
	queryName := reflect.TypeOf(query).Name()

	handler, exists := b.handlers[queryName]
	if !exists {
		return nil, fmt.Errorf("no handler registered for query %s", queryName)
	}

	return handler.Handle(query)
}

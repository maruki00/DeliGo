package pkgCqrs

import (
	"context"
	"fmt"
	"reflect"
)

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (_this *CommandBus) Register(command Command, handler CommandHandler) {
	commandName := reflect.TypeOf(command).Name()
	_this.handlers[commandName] = handler
}

func (_this *CommandBus) Dispatch(ctx context.Context, command Command) error {

	// commandName := reflect.TypeOf(command).Name()
	commandName := command.Name()

	if commandName == "" {
		return fmt.Errorf("the command is not registred")
	}

	handler, exists := _this.handlers[commandName]
	if !exists {
		return fmt.Errorf("no handler registered for command %s", commandName)
	}

	return handler.Handle(ctx, command)
}

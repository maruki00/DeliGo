package pkgCqrs

type CommandHandlerFunc func(Command) error
type Middleware func(next CommandHandler) CommandHandler

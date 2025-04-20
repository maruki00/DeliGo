package pkgCqrs

type Middleware func(next CommandHandler) CommandHandler
type CommandHandlerFunc func(Command) error

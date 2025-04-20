package pkgCqrs

type CommandHandler interface {
	Handle(command Command) error
}

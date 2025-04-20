package pkgCqrs

type Command interface {
	CommandName() string
}

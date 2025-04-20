package pkgCqrs

type QueryHandler interface {
	Handle(query Query) (interface{}, error)
}

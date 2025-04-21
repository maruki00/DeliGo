package userQueries

type FindUserByIdQuery struct {
	ID string
}

func (_this *FindUserByIdQuery) QueryName() string {
	return "FindUserByIdQuery"
}

package userQueries

type FindUserByIdQuery struct {
	ID string
}

func (_this *FindUserByIdQuery) CommandName() string {
	return "FindUserByIdQuery"
}

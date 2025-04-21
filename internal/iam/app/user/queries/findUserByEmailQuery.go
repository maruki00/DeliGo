package userQueries

type FindUserByEmailQuery struct {
	ID string
}

func (_this *FindUserByEmailQuery) CommandName() string {
	return "FindUserByEmailQuery"
}

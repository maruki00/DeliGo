package userQueries

type FindUserByEmailQuery struct {
	ID string
}

func (_this *FindUserByEmailQuery) QueryName() string {
	return "FindUserByEmailQuery"
}

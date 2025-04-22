package userQueries

type FindUserByEmailQuery struct {
	Email string
}

func (_this *FindUserByEmailQuery) QueryName() string {
	return "FindUserByEmailQuery"
}

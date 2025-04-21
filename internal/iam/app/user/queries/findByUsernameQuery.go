package userQueries

type FindUserByUsernameQuery struct {
	ID string
}

func (_this *FindUserByUsernameQuery) QueryName() string {
	return "FindUserByUsernameQuery"
}

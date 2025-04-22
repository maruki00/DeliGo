package userQueries

type FindUserByUsernameQuery struct {
	Username string
}

func (_this *FindUserByUsernameQuery) QueryName() string {
	return "FindUserByUsernameQuery"
}

package userQueries

type FindUserByUsernameQuery struct {
	Username string
}

func (_this *FindUserByUsernameQuery) Name() string {
	return "FindUserByUsernameQuery"
}

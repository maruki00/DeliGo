package userQueries

type FindUserByUsernameQuery struct {
	Key      string
	Username string
}

func (_this *FindUserByUsernameQuery) Name() string {
	return "FindUserByUsernameQuery"
}

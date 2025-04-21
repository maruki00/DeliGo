package userQueries

type FindUserByUsernameQuery struct {
	ID string
}

func (_this *FindUserByUsernameQuery) CommandName() string {
	return "FindUserByUsernameQuery"
}

package userQueries

type FindUserByEmailQuery struct {
	Email string
}

func (_this *FindUserByEmailQuery) Name() string {
	return "FindUserByEmailQuery"
}

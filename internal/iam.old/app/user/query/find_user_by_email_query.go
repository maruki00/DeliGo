package userQuery

type FindUserByEmailQuery struct {
	Key   string
	Email string
}

func (_this *FindUserByEmailQuery) Name() string {
	return "FindUserByEmailQuery"
}

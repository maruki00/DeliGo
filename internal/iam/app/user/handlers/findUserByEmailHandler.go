package userHandlers

type FindUserByEmailHandler struct {
	ID string
}

func (_this *FindUserByEmailHandler) Handle() string {
	return "FindUserByEmailQuery"
}

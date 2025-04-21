package userHandlers

type FindUserByUsernameHandler struct {
	ID string
}

func (_this *FindUserByUsernameHandler) Handle() string {
	return "FindUserByUsernameQuery"
}

package userHandlers

type DeleteUserHandler struct {
}

func (_this *DeleteUserHandler) Handle() string {
	return "DeleteUserCommand"
}

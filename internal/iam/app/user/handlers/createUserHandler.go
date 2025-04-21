package userHandlers

type CreateUserHandler struct {
}

func (_this *CreateUserHandler) Handle() string {
	return "CreateUserCommand"
}

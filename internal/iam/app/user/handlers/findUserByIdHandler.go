package userHandlers

type FindUserByIdHandler struct {
	ID string
}

func (_this *FindUserByIdHandler) Handle() string {
	return "FindUserByIdQuery"
}

package userHandlers

type ListUsersByTenantHandler struct {
	ID string
}

func (_this *ListUsersByTenantHandler) Handle() string {
	return "ListUsersByTenantQuery"
}

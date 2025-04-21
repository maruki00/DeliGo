package userHandlers

type ListUsersByTenantHandler struct {
	ID string
}

func (_this *ListUsersByTenantHandlery) Handle() string {
	return "ListUsersByTenantQuery"
}

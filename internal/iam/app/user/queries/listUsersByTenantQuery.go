package userQueries

type ListUsersByTenantQuery struct {
	ID string
}

func (_this *ListUsersByTenantQuery) CommandName() string {
	return "ListUsersByTenantQuery"
}

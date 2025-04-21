package userQueries

type ListUsersByTenantQuery struct {
	ID string
}

func (_this *ListUsersByTenantQuery) QueryName() string {
	return "ListUsersByTenantQuery"
}

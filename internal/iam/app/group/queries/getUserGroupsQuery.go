package userQueries

type GetUserGroupsQuery struct {
	ID string
}

func (_this *GetUserGroupsQuery) QueryName() string {
	return "GetUserGroupsQuery"
}

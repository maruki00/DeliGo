package models

type ShopEmployee struct {
	Id         int `json: "id"`
	ShopId     int `json: "shop_id"`
	EmployeeId int `json: "employee_id"`
}

func (obj *ShopEmployee) GetId() int {
	return obj.Id
}
func (obj *ShopEmployee) GetShopId() int {
	return obj.ShopId
}
func (obj *ShopEmployee) GetEmployeeId() int {
	return obj.EmployeeId
}

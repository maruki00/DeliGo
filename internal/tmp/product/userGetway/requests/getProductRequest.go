package requests

type GetProductRequest struct {
	Id int `validate: "required" josn:"id"`
}

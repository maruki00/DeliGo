package requests

type DeleteProductRequest struct {
	Id int `validate: "required" josn:"id"`
}

package requests

type MultipleProductstRequest struct {
	Ids string `validate: "required" josn:"ids"`
}

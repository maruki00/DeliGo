package requests

type SearchProductRequest struct {
	Query string `validate : "required" json:"query"`
}

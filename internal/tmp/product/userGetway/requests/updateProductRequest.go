package requests

type UpdateProductRequerst struct {
	Id    int     `validate:"required" json:"id"`
	Label string  `validate:"required" json:"label"`
	Price float32 `validate:"required" json:"price"`
	Type  string  `validate:"required" json:"type"`
}

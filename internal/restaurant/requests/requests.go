package requests

type CreateRestaurantRequest struct {
	Name    string `json:"name" binding:"required,max=255"`
	Address string `json:"address" binding:"required,max=255"`
}

type UpdateRestaurantStatusRequest struct {
	IsOpen *bool `json:"is_open" binding:"required"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	IsAvailable *bool   `json:"is_available" binding:"required"`
}

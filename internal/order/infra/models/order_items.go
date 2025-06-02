package models

import (
	shared_models "deligo/internal/shared/infra/models"
)

type OrderItems struct {
	shared_models.BaseModel
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Qty       int     `json:"qty"`
	UnitPrice float32 `json:"unit_price"`
}

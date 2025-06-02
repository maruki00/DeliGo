package models

import (
	shared_models "deligo/internal/shared/infra/models"
)

type Order struct {
	shared_models.BaseModel
	OrderFingerprint string  `json:"order_fingerprint"`
	CostumerId       int     `json:"costumer_id"`
	Cost             float32 `json:"cost"`
	Status           int     `json:"status"`
}

package shared_models

import (
	"database/sql"
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"time"
)

type DeletedAt sql.NullTime
type Model struct {
	ID        shared_valueobject.ID `gorm:"primarykey" json:"id"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt DeletedAt             `gorm:"index" json:"deleted_at"`
}

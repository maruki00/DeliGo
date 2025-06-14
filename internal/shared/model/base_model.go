package shared_model

import (
	"database/sql"
	shared_valueobject "github.com/maruki00/deligo/internal/shared/domain/valueObjects"
	"time"
)

type DeletedAt sql.NullTime
type BaseModel struct {
	ID        shared_valueobject.ID `gorm:"primarykey" json:"id"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt DeletedAt             `gorm:"index" json:"deleted_at"`
}

func (_this *BaseModel) GetID() shared_valueobject.ID {
	return _this.ID
}

func (_this *BaseModel) GetCreatedAt() time.Time {
	return _this.CreatedAt
}

func (_this *BaseModel) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

func (_this *BaseModel) GetDeletedAt() DeletedAt {
	return _this.DeletedAt
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"type:char(100);not null" json:"name"`
	CreateAt time.Time `json:"create_at"`
}

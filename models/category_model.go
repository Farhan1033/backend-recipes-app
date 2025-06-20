package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	CreateAt string    `json:"create_at"`
}

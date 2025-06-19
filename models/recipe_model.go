package models

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CategoryId  uuid.UUID `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryId" json:"-"`
	Title       string    `gorm:"type:char(255)" json:"title"`
	Description string    `grom:"text" json:"description"`
	Steps       string    `grom:"text" json:"steps"`
	ImageUrl    string    `grom:"char(255)" json:"image_url"`
	CreateAt    time.Time `json:"create_at"`
}

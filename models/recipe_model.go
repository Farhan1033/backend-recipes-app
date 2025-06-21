package models

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CategoryId  uuid.UUID `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryId" json:"category"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CookingTime int       `gorm:"type:int" json:"cooking_time"`
	Portion     int       `gorm:"type:int" json:"portion"`
	Steps       string    `gorm:"type:text" json:"steps"`
	ImageUrl    string    `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"created_at"`
}

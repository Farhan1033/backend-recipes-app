package models

import (
	"time"

	"github.com/google/uuid"
)

type RecipeIngredient struct {
	ID           uuid.UUID  `gorm:"type:uuid; primaryKey" json:"id"`
	RecipeId     uuid.UUID  `gorm:"type:uuid" json:"recipe_id"`
	IngredientId uuid.UUID  `gorm:"type:uuid" json:"ingredient_id"`
	Recipe       Recipe     `gorm:"foreignKey:RecipeId" json:"-"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientId" json:"-"`
	Quantity     float64    `gorm:"type:numeric(10,2)" json:"quantity"`
	Unit         string     `gorm:"type:varchar(50)" json:"unit"`
	CreatedAt    time.Time  `gorm:"autoCreateTime; autoUpdateTime" json:"created_at"`
}

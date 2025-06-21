package dto

import (
	"github.com/google/uuid"
)

type AddMultipleRecipeIngredientsRequest struct {
	RecipeID   uuid.UUID `json:"recipe_id" binding:"required"`
	Ingredients []struct {
		IngredientID uuid.UUID `json:"ingredient_id" binding:"required"`
		Quantity     int       `json:"quantity" binding:"required"`
		Unit         string    `json:"unit" binding:"required"`
	} `json:"ingredients" binding:"required,dive"`
}

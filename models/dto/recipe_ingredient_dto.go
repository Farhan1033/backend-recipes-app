package dto

import (
	"github.com/google/uuid"
)

type IngredientInRecipe struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity float64   `json:"quantity"`
	Unit     string    `json:"unit"`
}

type RecipeWithIngredients struct {
	ID          uuid.UUID            `json:"id"`
	RecipeId    uuid.UUID            `json:"recipe_id"`
	Title       string               `json:"title"`
	Category    string               `json:"category"`
	CookingTime int                  `json:"cooking_time"`
	Portion     int                  `json:"portion"`
	Description string               `json:"description"`
	Steps       string               `json:"steps"`
	ImageUrl    string               `json:"image_url"`
	CreatedAt   string               `json:"created_at"`
	Ingredients []IngredientInRecipe `json:"ingredients"`
}

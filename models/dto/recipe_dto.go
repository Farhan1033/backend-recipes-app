package dto

import "github.com/google/uuid"

type RecipeResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Steps       string    `json:"steps"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   string    `json:"created_at"`
}

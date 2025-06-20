package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"

	"github.com/google/uuid"
)

func CreateIngredient(ingredient *models.Ingredient) error {
	return config.DB.Create(ingredient).Error
}

func GetAllIngredient() ([]models.Ingredient, error) {
	var ingredient []models.Ingredient
	err := config.DB.Find(&ingredient).Error
	return ingredient, err
}

func DeleteIngredient(id uuid.UUID) error {
	return config.DB.Delete(&models.Ingredient{}, "id = ?", id).Error
}

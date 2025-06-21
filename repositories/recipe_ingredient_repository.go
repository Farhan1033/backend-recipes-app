package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"
)

func GetAllRecipeIngredient() ([]models.RecipeIngredient, error) {
	var recipeIngredient []models.RecipeIngredient
	err := config.DB.Preload("Recipe").Preload("Ingredient").Find(&recipeIngredient).Error
	return recipeIngredient, err
}

func CreateRecipeIngredient(recipeIngredient *models.RecipeIngredient) error {
	return config.DB.Create(recipeIngredient).Error
}

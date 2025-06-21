package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"

	"github.com/google/uuid"
)

func GetAllRecipeIngredient() ([]models.RecipeIngredient, error) {
	var recipeIngredient []models.RecipeIngredient
	err := config.DB.Preload("Recipe.Category").Preload("Ingredient").Find(&recipeIngredient).Error
	return recipeIngredient, err
}

func GetRecipeIngredientById(id uuid.UUID) (models.RecipeIngredient, error) {
	var recipeIngredient models.RecipeIngredient

	err := config.DB.
		Preload("Recipe.Category").
		Preload("Ingredient").
		First(&recipeIngredient, "recipe_id = ?", id).Error
	return recipeIngredient, err
}

func CreateRecipeIngredient(recipeIngredient *models.RecipeIngredient) error {
	return config.DB.Create(recipeIngredient).Error
}

func UpdateRecipeIngredient(id uuid.UUID, input *models.RecipeIngredient) error {
	var recipeIngredient models.RecipeIngredient

	if err := config.DB.First(&recipeIngredient, "id = ?", id).Error; err != nil {
		return err
	}

	recipeIngredient.RecipeId = input.RecipeId
	recipeIngredient.IngredientId = input.IngredientId
	recipeIngredient.Quantity = input.Quantity
	recipeIngredient.Unit = input.Unit

	return config.DB.Save(&recipeIngredient).Error
}

func DeleteRecipeIngredient(id uuid.UUID) error {
	return config.DB.Delete(&models.RecipeIngredient{}, "id = ?", id).Error
}

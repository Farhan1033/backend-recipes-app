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

func GetRecipeIngredientsByRecipeId(recipeId uuid.UUID) ([]models.RecipeIngredient, error) {
	var recipeIngredients []models.RecipeIngredient

	err := config.DB.
		Preload("Recipe.Category").
		Preload("Ingredient").
		Where("recipe_id = ?", recipeId).
		Find(&recipeIngredients).Error

	return recipeIngredients, err
}

func CreateRecipeIngredient(recipeIngredient []models.RecipeIngredient) error {
	if len(recipeIngredient) == 0 {
		return nil
	}

	return config.DB.Create(&recipeIngredient).Error
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

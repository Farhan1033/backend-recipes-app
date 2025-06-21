package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"

	"github.com/google/uuid"
)

func GetAllRecipe() ([]models.Recipe, error) {
	var recipe []models.Recipe
	err := config.DB.Preload("Category").Find(&recipe).Error
	return recipe, err
}

func GetRecipeById(id uuid.UUID) (models.Recipe, error) {
	var recipe models.Recipe
	err := config.DB.Preload("Category").First(&recipe, "id = ?", id).Error
	return recipe, err
}

func SearchRecipe(keyword string) ([]models.Recipe, error) {
	var recipe []models.Recipe

	if err := config.DB.Preload("Category").Where("LOWER(title) LIKE LOWER(?)", "%"+keyword+"%").Find(&recipe).Error; err != nil {
		return nil, err
	}

	return recipe, nil
}

func CreateRecipe(recipe *models.Recipe) error {
	return config.DB.Create(recipe).Error
}

func UpdateRecipe(id uuid.UUID, input *models.Recipe) error {
	var recipe models.Recipe

	if err := config.DB.First(&recipe, "id = ?", id).Error; err != nil {
		return err
	}

	recipe.Title = input.Title
	recipe.CategoryId = input.CategoryId
	recipe.Description = input.Description
	recipe.Steps = input.Steps
	recipe.ImageUrl = input.ImageUrl

	return config.DB.Save(&recipe).Error
}

func DeleteRecipe(id uuid.UUID) error {
	return config.DB.Delete(&models.Recipe{}, "id = ?", id).Error
}

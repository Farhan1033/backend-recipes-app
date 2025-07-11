package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"

	"github.com/google/uuid"
)

func CreateCategory(category *models.Category) error {
	return config.DB.Create(category).Error
}

func GetAllCategory() ([]models.Category, error) {
	var category []models.Category
	err := config.DB.Find(&category).Error
	return category, err
}

func DeleteCategory(id uuid.UUID) error {
	return config.DB.Delete(&models.Category{}, "id = ?", id).Error
}

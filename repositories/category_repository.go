package repositories

import (
	"backend-recipes/config"
	"backend-recipes/models"
	"time"

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

func GetCategoryById(id uuid.UUID) (models.Category, error) {
	var category models.Category
	if err := config.DB.First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func UpdateCategory(id uuid.UUID, input *models.Category) error {
	var category models.Category

	if err := config.DB.First(&category, "id = ?", id).Error; err != nil {
		return err
	}

	category.Name = input.Name
	category.CreateAt = time.Now()

	return config.DB.Save(&category).Error

}

func DeleteCategory(id uuid.UUID) error {
	return config.DB.Delete(&models.Category{}, "id = ?", id).Error
}

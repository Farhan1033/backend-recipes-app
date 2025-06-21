package controllers

import (
	"backend-recipes/models"
	"backend-recipes/models/dto"
	"backend-recipes/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllRecipeIngredient(c *gin.Context) {
	data, err := repositories.GetAllRecipeIngredient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grouped := make(map[uuid.UUID]*dto.RecipeWithIngredients)

	for _, ri := range data {
		recipeID := ri.Recipe.ID
		if _, found := grouped[recipeID]; !found {
			grouped[recipeID] = &dto.RecipeWithIngredients{
				ID:          ri.ID,
				RecipeId:    ri.Recipe.ID,
				Title:       ri.Recipe.Title,
				Category:    ri.Recipe.Category.Name,
				CookingTime: ri.Recipe.CookingTime,
				Portion:     ri.Recipe.Portion,
				Description: ri.Recipe.Description,
				Steps:       ri.Recipe.Steps,
				ImageUrl:    ri.Recipe.ImageUrl,
				CreatedAt:   ri.CreatedAt.Format("2006-01-02 15:04"),
				Ingredients: []dto.IngredientInRecipe{},
			}
		}

		grouped[recipeID].Ingredients = append(grouped[recipeID].Ingredients, dto.IngredientInRecipe{
			ID:       ri.Ingredient.ID,
			Name:     ri.Ingredient.Name,
			Quantity: float64(ri.Quantity),
			Unit:     ri.Unit,
		})
	}

	var response []dto.RecipeWithIngredients
	for _, v := range grouped {
		response = append(response, *v)
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data grouped by recipe",
		Data:    response,
	})
}

func GetRecipeIngredientById(c *gin.Context) {
	idParam := c.Param("id")

	recipeId, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	recipe, err := repositories.GetRecipeById(recipeId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	recipeIngredients, err := repositories.GetRecipeIngredientsByRecipeId(recipeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipe ingredients"})
		return
	}

	var ingredients []dto.IngredientInRecipe
	for _, ri := range recipeIngredients {
		ingredients = append(ingredients, dto.IngredientInRecipe{
			ID:       ri.Ingredient.ID,
			Name:     ri.Ingredient.Name,
			Quantity: ri.Quantity,
			Unit:     ri.Unit,
		})
	}

	result := dto.RecipeWithIngredients{
		ID:          recipe.ID,
		RecipeId:    recipe.ID,
		Title:       recipe.Title,
		Category:    recipe.Category.Name,
		CookingTime: recipe.CookingTime,
		Portion:     recipe.Portion,
		Description: recipe.Description,
		Steps:       recipe.Steps,
		ImageUrl:    recipe.ImageUrl,
		CreatedAt:   recipe.CreatedAt.Format("2006-01-02 15:04"),
		Ingredients: ingredients,
	}

	c.JSON(http.StatusOK, result)
}


func CreateRecipeIngredient(c *gin.Context) {
	var input dto.AddMultipleRecipeIngredientsRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data []models.RecipeIngredient
	for _, item := range input.Ingredients {
		data = append(data, models.RecipeIngredient{
			ID:           uuid.New(),
			RecipeId:     input.RecipeID,
			IngredientId: item.IngredientID,
			Quantity:     float64(item.Quantity),
			Unit:         item.Unit,
			CreatedAt:    time.Now(),
		})
	}

	if err := repositories.CreateRecipeIngredient(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambahkan ingredients",
	})
}

func UpdateRecipeIngredient(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var input models.RecipeIngredient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateRecipeIngredient(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil update recipe ingredient"})
}

func DeleteRecipeIngredient(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := repositories.DeleteRecipeIngredient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil hapus recipe ingredient"})
}

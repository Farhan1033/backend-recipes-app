package controllers

import (
	"backend-recipes/models"
	"backend-recipes/models/dto"
	"backend-recipes/repositories"
	"net/http"

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
				ID:          recipeID,
				Title:       ri.Recipe.Title,
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

func CreateRecipeIngredient(c *gin.Context) {
	var recipeIngredient models.RecipeIngredient

	if err := c.ShouldBindJSON(&recipeIngredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	recipeIngredient.ID = uuid.New()

	if err := repositories.CreateRecipeIngredient(&recipeIngredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil membuat data",
		Data:    recipeIngredient,
	})
}

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

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	data, err := repositories.GetRecipeIngredientById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe Ingredient not found"})
		return
	}

	recipe, err := repositories.GetRecipeById(data.Recipe.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	result := dto.RecipeWithIngredients{
		ID:          data.ID,
		RecipeId:    data.Recipe.ID,
		Title:       data.Recipe.Title,
		Category:    recipe.Category.Name,
		CookingTime: data.Recipe.CookingTime,
		Portion:     data.Recipe.Portion,
		Description: data.Recipe.Description,
		Steps:       data.Recipe.Steps,
		ImageUrl:    data.Recipe.ImageUrl,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04"),
		Ingredients: []dto.IngredientInRecipe{
			{
				ID:       data.Ingredient.ID,
				Name:     data.Ingredient.Name,
				Quantity: data.Quantity,
				Unit:     data.Unit,
			},
		},
	}

	c.JSON(http.StatusOK, result)
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

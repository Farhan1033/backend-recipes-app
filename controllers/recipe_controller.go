package controllers

import (
	"backend-recipes/models"
	"backend-recipes/models/dto"
	"backend-recipes/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllRecipe(c *gin.Context) {
	recipe, err := repositories.GetAllRecipe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	var response []dto.RecipeResponse

	for _, r := range recipe {
		i := dto.RecipeResponse{
			ID:          r.ID,
			Title:       r.Title,
			Category:    r.Category.Name,
			CookingTime: r.CookingTime,
			Portion:     r.Portion,
			Description: r.Description,
			Steps:       r.Steps,
			ImageUrl:    r.ImageUrl,
			CreatedAt:   r.CreatedAt.Format("2006-01-02 15:04"),
		}
		response = append(response, i)
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    response,
	})
}

func GetRecipeById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID Format"})
		return
	}

	recipe, err := repositories.GetRecipeById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.RecipeResponse{
		ID:          recipe.ID,
		Title:       recipe.Title,
		Category:    recipe.Category.Name,
		CookingTime: recipe.CookingTime,
		Portion:     recipe.Portion,
		Description: recipe.Description,
		Steps:       recipe.Steps,
		ImageUrl:    recipe.ImageUrl,
		CreatedAt:   recipe.CreatedAt.Format("2006-01-02 15:04"),
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    response,
	})
}

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	recipe.ID = uuid.New()

	if err := repositories.CreateRecipe(&recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil membuat data",
		Data:    recipe,
	})
}

func SearchRecipe(c *gin.Context) {
	keyword := c.Query("t")

	recipe, err := repositories.SearchRecipe(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.RecipeResponse

	for _, r := range recipe {
		i := dto.RecipeResponse{
			ID:          r.ID,
			Title:       r.Title,
			Category:    r.Category.Name,
			CookingTime: r.CookingTime,
			Portion:     r.Portion,
			Description: r.Description,
			Steps:       r.Steps,
			ImageUrl:    r.ImageUrl,
			CreatedAt:   r.CreatedAt.Format("2006-01-02 15:04"),
		}
		response = append(response, i)
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    response,
	})
}

func UpdateRecipe(c *gin.Context) {
	var recipe models.Recipe

	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID Format"})
		return
	}

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateRecipe(id, &recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil update data",
		Data:    recipe,
	})
}

func DeleteRecipe(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID Format"})
		return
	}

	if err != repositories.DeleteRecipe(id) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menghapus data",
	})
}

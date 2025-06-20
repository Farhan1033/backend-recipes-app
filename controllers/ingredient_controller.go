package controllers

import (
	"backend-recipes/models"
	"backend-recipes/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateIngredient(c *gin.Context) {
	var ingredient models.Ingredient

	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ingredient.ID = uuid.New()

	if err := repositories.CreateIngredient(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusCreated, Response{
		Message: "Berhasil membuat data",
		Data:    ingredient,
	})
}

func GetAllIngredient(c *gin.Context) {
	data, err := repositories.GetAllIngredient()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil mengambil data",
		Data:    data,
	})
}

func DeleteIngredient(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID Format",
		})
		return
	}

	if err := repositories.DeleteIngredient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	type Response struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menghapus data",
	})
}

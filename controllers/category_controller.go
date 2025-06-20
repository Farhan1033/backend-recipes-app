package controllers

import (
	"backend-recipes/models"
	"backend-recipes/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllCategory(c *gin.Context) {
	data, err := repositories.GetAllCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    data,
	})
}

func GetCategoryById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID Format",
		})
		return
	}

	data, errs := repositories.GetCategoryById(id)
	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errs.Error(),
		})
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    data,
	})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	category.ID = uuid.New()
	category.CreateAt = time.Now().Format(time.RFC3339)

	if err := repositories.CreateCategory(&category); err != nil {
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
		Data:    category,
	})
}

func DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID Format",
		})
		return
	}

	if err := repositories.DeleteCategory(id); err != nil {
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

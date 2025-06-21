package routes

import (
	"backend-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeIngredientRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/recipe-ingredient/", controllers.GetAllRecipeIngredient)
		api.POST("/recipe-ingredient/create", controllers.CreateRecipeIngredient)
	}
}

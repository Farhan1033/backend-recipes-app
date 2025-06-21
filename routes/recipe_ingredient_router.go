package routes

import (
	"backend-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeIngredientRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/recipe-ingredient/", controllers.GetAllRecipeIngredient)
		api.GET("/recipe-ingredient/:id", controllers.GetRecipeIngredientById)
		api.POST("/recipe-ingredient/create", controllers.CreateRecipeIngredient)
		api.PUT("/recipe-ingredient/update/:id", controllers.UpdateRecipeIngredient)
		api.DELETE("/recipe-ingredient/delete/:id", controllers.DeleteRecipeIngredient)
	}
}

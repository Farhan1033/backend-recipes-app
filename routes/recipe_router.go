package routes

import (
	"backend-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/recipes/", controllers.GetAllRecipe)
		api.GET("/recipes/:id", controllers.GetRecipeById)
		api.GET("/recipes/search", controllers.SearchRecipe)
		api.POST("/recipes/create", controllers.CreateRecipe)
		api.PUT("/recipes/update/:id", controllers.UpdateRecipe)
		api.DELETE("/recipes/delete/:id", controllers.DeleteRecipe)
	}
}

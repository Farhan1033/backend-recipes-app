package routes

import (
	"backend-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func IngredientRoute(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ingredient/", controllers.GetAllIngredient)
		api.POST("/ingredient/create", controllers.CreateIngredient)
		api.DELETE("/ingredient/delete/:id", controllers.DeleteIngredient)
	}
}

package routes

import (
	"backend-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/category/", controllers.GetAllCategory)
		api.POST("/category/create", controllers.CreateCategory)
		api.DELETE("/category/delete/:id", controllers.DeleteCategory)
	}
}

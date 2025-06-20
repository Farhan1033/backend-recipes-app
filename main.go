package main

import (
	"backend-recipes/config"
	"backend-recipes/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"0.0.0.0"}); err != nil {
		panic(fmt.Sprintf("Gagal set trusted proxy: %v", err))
	}

	r.GET("/kaithhealthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	routes.CategoryRoute(r)
	routes.IngredientRoute(r)

	port := os.Getenv("PORT")

	fmt.Printf("Server berjalan di http://0.0.0.0:%s\n", port)

	r.Run("0.0.0.0:" + port)
}

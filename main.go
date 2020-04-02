package main

import (
	"github.com/Bryce-Soghigian/databases-w-go/controllers"
	"github.com/Bryce-Soghigian/databases-w-go/models"
	"github.com/gin-gonic/gin"
)

// func get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"status":"up"}`))
// }

func main() {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.Run()
}

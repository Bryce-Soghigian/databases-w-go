package main

import (
	"github.com/Bryce-Soghigian/databases-w-go/controllers"
	"github.com/Bryce-Soghigian/databases-w-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBookById)
	r.PATCH("/books/:id", controllers.PatchBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()

}

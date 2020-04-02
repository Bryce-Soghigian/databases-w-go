package main

import (
	"net/http"

	"github.com/bryce-soghigian/databases-w-go/models"
	"github.com/gin-gonic/gin"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

func main() {
	r := gin.Default()
	db := models.SetupModels() // new

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Run()
}

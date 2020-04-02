package controllers

import (
	"net/http"

	"github.com/Bryce-Soghigian/databases-w-go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Author, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
func FindBookById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//fetch book model if it exists
	var book models.Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not in there boi"})
		return
	}
	c.JSON((http.StatusOK), gin.H{"data": book})
}

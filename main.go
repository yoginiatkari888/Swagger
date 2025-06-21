package main

import (
	"net/http"
	"strconv"

	_ "swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Model
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Go Programming", Author: "John Doe"},
	{ID: 2, Title: "REST APIs with Gin", Author: "Jane Doe"},
}

var nextID = 3

// @Summary Greet user
// @Description Returns a welcome message
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /books [post]
func createBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// Read All
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// Read by ID
func getBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// Update
func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, b := range books {
		if b.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// Delete
func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

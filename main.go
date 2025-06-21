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

// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body Book true "Book to create"
// @Success 201 {object} Book
// @Failure 400 {object} map[string]string
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

// @Summary Get all books
// @Description Retrieve the list of books
// @Tags books
// @Produce json
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// @Summary Get book by ID
// @Description Retrieve a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
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

// @Summary Welcome
// @Description Welcome message for the Book API
// @Tags root
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func weclome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to the Book API. Visit /swagger/index.html for docs."})
}

// @Summary Update a book
// @Description Update a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body Book true "Updated Book"
// @Success 200 {object} Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [put]
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

// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
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

// @title Book API
// @version 1.0
// @description A simple CRUD API built with Go and Gin
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", weclome)
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

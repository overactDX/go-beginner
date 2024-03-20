package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In the name of love", Author: "Cartoons", Quantity: 2},
	{ID: "2", Title: "Another Book", Author: "Another Author", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newBooks Book

	if err := c.BindJSON(&newBooks); err != nil {
		return
	}

	books = append(books, newBooks)
	c.IndentedJSON(http.StatusCreated, newBooks)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBooksByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func getBooksByID(c *gin.Context, id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}
func main() {
	route := gin.Default()

	route.GET("/books", getBooks)
	route.GET("/books/:id", bookById)
	route.POST("/books", createBooks)
	route.Run("localhost: 1000")
}
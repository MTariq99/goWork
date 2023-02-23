package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Auther string `json:"auther"`
}

var books = []Book{
	{ID: "1", Title: "harry porter", Auther: "J. K. Rowling"},
	{ID: "2", Title: "witcher", Auther: "zinzhu"},
	{ID: "3", Title: "Vikings", Auther: "floki"},
}

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			break
		}
	}
	c.Status(http.StatusNotFound)
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBookById(c *gin.Context) {
	bookId := c.Param("id")
	var updateBook Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	for i := range books {
		if books[i].ID == bookId {
			books[i] = updateBook
			c.JSON(http.StatusCreated, gin.H{
				"message": "book updated successfully",
			})
			return
		}
	}
	c.Status(http.StatusCreated)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("id")

	// for i := 0; i < len(books); i++ {
	// 	if books[i].ID == bookId {
	// 		for j := i; j < len(books); j++ {
	// 			books[j] = books[j+1]
	// 		}

	// 	}
	// }
	for i := range books {
		if books[i].ID == bookId {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)

}

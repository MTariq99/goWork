package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/books", listBookHandler)
	router.POST("/books", createBookHandler)
	router.DELETE("/books/:id", deleteBookHandler)
	// router := gin.New()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world!",
	// 	})
	// })
	// router.GET("/books", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, books)
	// })
	// router.POST("/books", func(c *gin.Context) {
	// 	var book Book
	// 	if err := c.ShouldBindJSON(&book); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"ERROR": err.Error(),
	// 		})
	// 	}
	// 	books = append(books, book)
	// 	c.JSON(http.StatusOK, book)
	// })
	// router.DELETE("books/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	for i, a := range books {
	// 		if a.ID == id {
	// 			books = append(books[:i], books[i+1:]...)
	// 			break
	// 		}
	// 	}
	// 	c.Status(http.StatusNoContent)

	// })
	fmt.Println("the server is Running on Port:http://localhost:8080/ ")
	router.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Auther string `json:"auther"`
}

var books = []Book{
	{ID: "1", Title: "Harry Porter", Auther: "J. K. Rowling"},
	{ID: "2", Title: "The Lord Of The Rings", Auther: "j.R.R.Tolkien"},
	{ID: "3", Title: "the wizard of OZ", Auther: "L.Frank Baum"},
}

func listBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)

}
func createBookHandler(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}

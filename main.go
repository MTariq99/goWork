package main

import (
	"bookstoreApi/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	bookRoutes := r.Group("/books")
	routes.RegisterBookStoreRoutes(bookRoutes)
	r.Run()
	fmt.Println("the server is running\n ")
}

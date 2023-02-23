package main

import (
	"bookstoreApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	bookRoutes := r.Group("/books")
	routes.RegisterBookStoreRoutes(bookRoutes)
	r.Run()
}

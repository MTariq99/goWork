package routes

import (
	"bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBookStoreRoutes(route *gin.RouterGroup) {
	route.GET("", controllers.GetAllBooks)
	route.GET("/:id", controllers.GetBookById)
	route.POST("", controllers.CreateBook)
	route.PUT("/:id", controllers.UpdateBookById)
	route.DELETE("/:id", controllers.DeleteBook)
}

package router

import (
	"books-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Books Routes
	r.GET("/books", controllers.GetBooks)
	r.GET("books/:id", controllers.GetBooks)
	r.POST("/books", controllers.GetBooks)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}

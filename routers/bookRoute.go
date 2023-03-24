package routers

import (
	"challange-06/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/books", controllers.GetAllBook)
	router.GET("/book/:id", controllers.GetBook)
	router.POST("/create-book", controllers.CreateBook)
	router.PUT("/update-book/:id", controllers.UpdateBook)
	router.DELETE("/delete-book/:id", controllers.DeleteBook)
	return router
}

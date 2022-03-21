package main

import (
	controller "pustaka-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controller.GetAllBooks)
	router.DELETE("/books/:id", controller.DeleteBook)
	router.POST("/books", controller.InsertBook)

	router.Run(":8080")
}

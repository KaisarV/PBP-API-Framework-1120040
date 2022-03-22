package main

import (
	controller "pustaka-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Book
	router.GET("/books", controller.GetAllBooks)
	router.DELETE("/books/:id", controller.DeleteBook)
	router.POST("/books", controller.InsertBook)
	router.PUT("/books/:id", controller.UpdateBooks)

	//User
	router.GET("/users", controller.GetAllUsers)

	router.Run(":8080")
}

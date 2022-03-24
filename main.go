package main

import (
	controller "pustaka-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/login", controller.UserLogin)
	router.GET("/logout", controller.Logout)

	router.GET("/books", controller.GetAllBooks)
	router.DELETE("/books/:id", controller.DeleteBook)
	router.POST("/books", controller.InsertBook)
	router.PUT("/books/:id", controller.UpdateBooks)

	//User
	router.GET("/users", controller.GetAllUsers)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.POST("/users", controller.InsertUser)
	router.PUT("/users/:id", controller.UpdateUsers)

	//Lendings
	router.GET("/lendings", controller.GetAllLending)
	router.DELETE("/lendings/:id", controller.DeleteLending)
	router.POST("/lendings", controller.InsertLendings)
	router.PUT("lendings/:id", controller.UpdateLendings)

	router.Run(":8080")
}

package main

import (
	controller "pustaka-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", controller.UserLogin)
	router.GET("/logout", controller.Logout)

	//1. Admin 2. Basic User
	admin := router.Group("/")
	admin.Use(controller.Authenticate(1))
	{
		admin.POST("/books", controller.InsertBook)
		admin.PUT("/books/:id", controller.UpdateBooks)
		admin.POST("/users", controller.InsertUser)

		//Lending
		admin.GET("/lendings", controller.GetAllLending)
		admin.DELETE("/lendings/:id", controller.DeleteLending)
		admin.POST("/lendings", controller.InsertLendings)
		admin.PUT("lendings/:id", controller.UpdateLendings)
	}

	router.GET("/books", controller.GetAllBooks)
	router.DELETE("/books/:id", controller.DeleteBook)

	//User
	router.GET("/users", controller.GetAllUsers)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.PUT("/users/:id", controller.UpdateUsers)

	router.Run(":8080")
}

package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "pustaka-api/models"
)

func GetAllUsers(c *gin.Context) {
	db := connect()
	var response model.UsersResponse
	defer db.Close()

	query := "SELECT Id, Name, Address, Email, UserType FROM users"
	id := c.Query("id")
	if id != "" {
		query += " WHERE id = " + id
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Email, &user.UserType); err != nil {
			log.Println(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = users
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	c.JSON(response.Status, response)
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	userId := c.Param("id")

	query, errQuery := db.Exec(`DELETE FROM users WHERE Id = ?;`, userId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "User not found"
		c.JSON(400, response)
		return
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Delete Data"
	} else {
		response.Status = 400
		response.Message = "Error Delete Data"
		log.Println(errQuery.Error())
	}

	c.JSON(response.Status, response)
}

func InsertUser(c *gin.Context) {

	db := connect()

	var user model.User
	var response model.UserResponse

	user.Name = c.PostForm("author")
	user.Address = c.PostForm("desc")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	user.UserType, _ = strconv.Atoi(c.PostForm("usertype"))

	if user.Name == "" {
		response.Status = 400
		response.Message = "Please Insert User's Name"
		c.JSON(response.Status, response)
		return
	}

	if user.Address == "" {
		response.Status = 400
		response.Message = "Please Insert User's Address"
		c.JSON(response.Status, response)
		return
	}

	if user.Email == "" {
		response.Status = 400
		response.Message = "Please Insert User's Email"
		c.JSON(response.Status, response)
		return
	}

	if user.Password == "" {
		response.Status = 400
		response.Message = "Please Insert User's Password"
		c.JSON(response.Status, response)
		return
	}

	if user.UserType == 0 {
		response.Status = 400
		response.Message = "Please Insert User's type"
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO users(Name, Address, Email, Password, UserType) VALUES(?, ?, ?, ?, ?)", user.Name, user.Address, user.Email, user.Password, user.UserType)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		user.ID = int(id)
		response.Data = user
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}

	c.JSON(response.Status, response)
}

// func UpdateBooks(c *gin.Context) {
// 	db := connect()

// 	var book model.Book
// 	var response model.BookResponse

// 	bookId := c.Param("id")
// 	book.Title = c.PostForm("title")
// 	book.Author = c.PostForm("author")
// 	book.Description = c.PostForm("desc")
// 	book.Price, _ = strconv.Atoi(c.PostForm("price"))
// 	book.Rating, _ = strconv.Atoi(c.PostForm("rating"))

// 	rows, _ := db.Query("SELECT * FROM books WHERE id = ?", bookId)
// 	var prevDatas []model.Book
// 	var prevData model.Book

// 	for rows.Next() {
// 		if err := rows.Scan(&prevData.ID, &prevData.Title, &prevData.Author, &prevData.Description, &prevData.Price, &prevData.Rating); err != nil {
// 			log.Println(err.Error())
// 		} else {
// 			prevDatas = append(prevDatas, prevData)
// 		}
// 	}

// 	if len(prevDatas) > 0 {
// 		if book.Title == "" {
// 			book.Title = prevDatas[0].Title
// 		}
// 		if book.Author == "" {
// 			book.Author = prevDatas[0].Author
// 		}
// 		if book.Description == "" {
// 			book.Description = prevDatas[0].Description
// 		}
// 		if book.Price == 0 {
// 			book.Price = prevDatas[0].Price
// 		}
// 		if book.Rating == 0 {
// 			book.Rating = prevDatas[0].Rating
// 		}

// 		_, errQuery := db.Exec(`UPDATE books SET Title = ?, Author = ?, Description = ?, Price = ?, Rating = ? WHERE id = ?`, book.Title, book.Author, book.Description, book.Price, book.Rating, bookId)

// 		if errQuery == nil {
// 			response.Status = 200
// 			response.Message = "Success Update Data"
// 			id, _ := strconv.Atoi(bookId)
// 			book.ID = id
// 			response.Data = book

// 		} else {
// 			response.Status = 400
// 			response.Message = "Error Update Data"

// 			log.Println(errQuery)
// 		}
// 	} else {
// 		response.Status = 400
// 		response.Message = "Data Not Found"
// 	}

// 	c.JSON(response.Status, response)
// }

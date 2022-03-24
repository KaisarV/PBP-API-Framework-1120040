package controllers

import (
	"log"
	"net/http"
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

	user.Name = c.PostForm("name")
	user.Address = c.PostForm("address")
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

func UpdateUsers(c *gin.Context) {
	db := connect()

	var user model.User
	var response model.UserResponse

	userId := c.Param("id")
	user.Name = c.PostForm("name")
	user.Address = c.PostForm("address")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	user.UserType, _ = strconv.Atoi(c.PostForm("usertype"))

	rows, _ := db.Query("SELECT * FROM users WHERE id = ?", userId)
	var prevDatas []model.User
	var prevData model.User

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Name, &prevData.Address, &prevData.Email, &prevData.Password, &prevData.UserType); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if user.Name == "" {
			user.Name = prevDatas[0].Name
		}
		if user.Address == "" {
			user.Address = prevDatas[0].Address
		}
		if user.Email == "" {
			user.Email = prevDatas[0].Email
		}
		if user.Password == "" {
			user.Password = prevDatas[0].Password
		}
		if user.UserType == 0 {
			user.UserType = prevDatas[0].UserType
		}

		_, errQuery := db.Exec(`UPDATE users SET Name = ?, Address = ?, Email = ?, Password = ?, UserType = ? WHERE id = ?`, user.Name, user.Address, user.Email, user.Password, user.UserType, userId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			id, _ := strconv.Atoi(userId)
			user.ID = id
			response.Data = user
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	c.JSON(response.Status, response)
}

func UserLogin(c *gin.Context) {
	db := connect()
	defer db.Close()

	name := c.PostForm("name")
	password := c.PostForm("password")

	rows, err := db.Query("SELECT * FROM users WHERE name=? AND password=?",
		name,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) == 1 {
		generateToken(c, user.ID, user.Name, user.UserType)
		sendSuccessResponse(c)
	} else {
		sendErrorResponse(c)
	}
}

func Logout(c *gin.Context) {
	resetUserToken(c)
	sendSuccessResponse(c)
}

func sendSuccessResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 200
	response.Message = "Success"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func sendErrorResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusBadRequest, response)
}

func sendUnAuthorizedResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusUnauthorized, response)
}

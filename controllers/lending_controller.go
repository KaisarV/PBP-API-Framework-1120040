package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "pustaka-api/models"
)

func GetAllLending(c *gin.Context) {
	db := connect()
	var response model.LendingsResponse
	defer db.Close()

	query := "SELECT l.LendId, b.Id, b.Title, b.Author, b.Description, b.Price, b.Rating, u.Id, u.Name, u.Address, u.Email, u.Password, u.UserType, l.LendingDate FROM users u JOIN lendings l ON u.Id = l.UserId JOIN books b ON l.BookId = b.Id"

	bookId := c.Query("bookId")
	if bookId != "" {
		query += " WHERE l.BookId = " + bookId
	}

	userId := c.Query("userId")
	if userId != "" {
		query += " WHERE UserId = " + userId
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var book model.Book
	var user model.User
	var lendings []model.Lending
	var lending model.Lending

	for rows.Next() {
		if err := rows.Scan(&lending.LendId, &book.ID, &book.Title, &book.Author, &book.Description, &book.Price, &book.Rating, &user.ID, &user.Name, &user.Address, &user.Email, &user.Password, &user.UserType, &lending.LendingDate); err != nil {
			log.Println(err.Error())
		} else {
			lending.User = user
			lending.Book = book
			lendings = append(lendings, lending)
		}
	}

	if len(lendings) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = lendings
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func DeleteLending(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	lendingId := c.Param("id")

	query, errQuery := db.Exec(`DELETE FROM lendings WHERE LendId = ?;`, lendingId)

	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "Lending not found"
		c.Header("Content-Type", "application/json")
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
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func InsertLendings(c *gin.Context) {

	db := connect()

	var lending model.LendingIndex
	var response model.LendingIndexResponse

	lending.BookId, _ = strconv.Atoi(c.PostForm("BookId"))
	lending.UserId, _ = strconv.Atoi(c.PostForm("UserId"))

	if lending.BookId == 0 {
		response.Status = 400
		response.Message = "Please Insert Book"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if lending.UserId == 0 {
		response.Status = 400
		response.Message = "Please Insert User"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO lendings(BookId, UserId) VALUES(?, ?)", lending.BookId, lending.UserId)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		lending.LendId = int(id)
		response.Data = lending
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func UpdateLendings(c *gin.Context) {
	db := connect()

	var lending model.LendingIndex
	var response model.LendingIndexResponse

	lendingId := c.Param("id")
	lending.BookId, _ = strconv.Atoi(c.PostForm("BookId"))
	lending.UserId, _ = strconv.Atoi(c.PostForm("UserId"))

	rows, _ := db.Query("SELECT * FROM lendings WHERE LendId = ?", lendingId)
	var prevDatas []model.LendingIndex
	var prevData model.LendingIndex

	for rows.Next() {
		if err := rows.Scan(&prevData.LendId, &prevData.BookId, &prevData.UserId, &prevData.LendingDate); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if lending.BookId == 0 {
			lending.BookId = prevDatas[0].LendId
		}
		if lending.UserId == 0 {
			lending.UserId = prevDatas[0].UserId
		}

		_, errQuery := db.Exec(`UPDATE lendings SET BookId = ?, UserId = ? WHERE LendId = ?`, lending.BookId, lending.UserId, lendingId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			id, _ := strconv.Atoi(lendingId)
			lending.LendId = id
			response.Data = lending
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

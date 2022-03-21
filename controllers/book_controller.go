package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "pustaka-api/models"
)

func GetAllBooks(c *gin.Context) {

	db := connect()
	var response model.BooksResponse
	defer db.Close()

	query := "SELECT * FROM books"
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

	var book model.Book
	var books []model.Book

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price, &book.Rating); err != nil {
			log.Println(err.Error())
		} else {
			books = append(books, book)
		}
	}

	if len(books) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = books
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	c.JSON(response.Status, response)
}

func DeleteBook(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	bookId := c.Param("id")

	query, errQuery := db.Exec(`DELETE FROM books WHERE Id = ?;`, bookId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "Book not found"
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

func InsertBook(c *gin.Context) {

	db := connect()

	var book model.Book
	var response model.BookResponse

	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Description = c.PostForm("desc")
	book.Price, _ = strconv.Atoi(c.PostForm("price"))
	book.Rating, _ = strconv.Atoi(c.PostForm("rating"))

	if book.Title == "" {
		response.Status = 400
		response.Message = "Please Insert Book Title"
		c.JSON(response.Status, response)
		return
	}

	if book.Author == "" {
		response.Status = 400
		response.Message = "Please Insert Book Author"
		c.JSON(response.Status, response)
		return
	}

	if book.Description == "" {
		response.Status = 400
		response.Message = "Please Insert Book Description"
		c.JSON(response.Status, response)
		return
	}

	if book.Price == 0 {
		response.Status = 400
		response.Message = "Please Insert Book Price"
		c.JSON(response.Status, response)
		return
	}

	if book.Rating == 0 {
		response.Status = 400
		response.Message = "Please Insert Book Rating"
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO books(Title, Author, Description, Price, Rating) VALUES(?, ?, ?, ?, ?)", book.Title, book.Author, book.Description, book.Price, book.Rating)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		book.ID = int(id)
		response.Data = book
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}

	c.JSON(response.Status, response)
}

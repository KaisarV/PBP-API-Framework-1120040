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
		c.Header("Content-Type", "application/json")
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
	c.Header("Content-Type", "application/json")
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
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if book.Author == "" {
		response.Status = 400
		response.Message = "Please Insert Book Author"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if book.Description == "" {
		response.Status = 400
		response.Message = "Please Insert Book Description"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if book.Price == 0 {
		response.Status = 400
		response.Message = "Please Insert Book Price"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if book.Rating == 0 {
		response.Status = 400
		response.Message = "Please Insert Book Rating"
		c.Header("Content-Type", "application/json")
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
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func UpdateBooks(c *gin.Context) {
	db := connect()

	var book model.Book
	var response model.BookResponse

	bookId := c.Param("id")
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Description = c.PostForm("desc")
	book.Price, _ = strconv.Atoi(c.PostForm("price"))
	book.Rating, _ = strconv.Atoi(c.PostForm("rating"))

	rows, _ := db.Query("SELECT * FROM books WHERE id = ?", bookId)
	var prevDatas []model.Book
	var prevData model.Book

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Title, &prevData.Author, &prevData.Description, &prevData.Price, &prevData.Rating); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if book.Title == "" {
			book.Title = prevDatas[0].Title
		}
		if book.Author == "" {
			book.Author = prevDatas[0].Author
		}
		if book.Description == "" {
			book.Description = prevDatas[0].Description
		}
		if book.Price == 0 {
			book.Price = prevDatas[0].Price
		}
		if book.Rating == 0 {
			book.Rating = prevDatas[0].Rating
		}

		_, errQuery := db.Exec(`UPDATE books SET Title = ?, Author = ?, Description = ?, Price = ?, Rating = ? WHERE id = ?`, book.Title, book.Author, book.Description, book.Price, book.Rating, bookId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			id, _ := strconv.Atoi(bookId)
			book.ID = id
			response.Data = book

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

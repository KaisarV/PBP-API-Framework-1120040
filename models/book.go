package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type BooksResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Book `json:"data,omitempty"`
}

type BookResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Book   `json:"data,omitempty"`
}

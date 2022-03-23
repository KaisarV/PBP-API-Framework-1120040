package models

type Lending struct {
	LendId      int    `json:"lendId"`
	Book        Book   `json:"book"`
	User        User   `json:"user"`
	LendingDate string `json:"lendingDate,omitempty"`
}

type LendingIndex struct {
	LendId      int    `json:"lendId"`
	BookId      int    `json:"book"`
	UserId      int    `json:"user"`
	LendingDate string `json:"lendingDate,omitempty"`
}

type LendingIndexResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    LendingIndex `json:"data,omitempty"`
}
type LendingsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Lending `json:"data,omitempty"`
}

type LendingResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Lending `json:"data,omitempty"`
}

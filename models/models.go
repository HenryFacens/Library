package models

import "github.com/gocql/gocql"

type Book struct {
	ID          gocql.UUID `json:"id"`
	Author      string     `json:"author"`
	Genre       string     `json:"genre"`
	PublishYear int        `json:"publish_year"`
	Title       string     `json:"title"`
}

// Definir estrutura para capturar o body JSON
type RequestBody struct {
	UserID string `json:"id"`
}
type User struct {
	ID    gocql.UUID `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
}

type Borrow struct {
	BookID     gocql.UUID `json:"book_id"`
	UserID     gocql.UUID `json:"user_id"`
	BorrowDate string     `json:"borrow_date"`
}

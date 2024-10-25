package main

import (
	"library-system/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.InitCassandraSession("127.0.0.1")

	http.HandleFunc("/admin/book", handlers.AddBook)           // Adicionar livro
	http.HandleFunc("/admin/books", handlers.GetBooks)         // Listar livros
	http.HandleFunc("/admin/book/delete", handlers.DeleteBook) // Remover livro

	http.HandleFunc("/admin/user", handlers.AddUser)                         // Adicionar usuário
	http.HandleFunc("/admin/users", handlers.GetUsers)                       // Listar usuários
	http.HandleFunc("/admin/user/delete", handlers.DeleteUser)               // Remover usuário
	http.HandleFunc("/admin/user/recomendation", handlers.GetRecommendation) // Remover usuário

	http.HandleFunc("/student/borrow", handlers.BorrowBook)         // Emprestar livro
	http.HandleFunc("/student/borrowed", handlers.GetBorrowedBooks) // Listar empréstimos
	http.HandleFunc("/student/return", handlers.ReturnBook)         // Devolver livro

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

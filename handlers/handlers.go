package handlers

import (
	"encoding/json"
	"library-system/models"
	"log"
	"net/http"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func InitCassandraSession(clusterIP string) {
	cluster := gocql.NewCluster(clusterIP)
	cluster.Keyspace = "library"
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Erro ao conectar ao Cassandra: ", err)
	}
	log.Println("Conexão com Cassandra estabelecida")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	book.ID = gocql.TimeUUID()
	err := session.Query(`
        INSERT INTO books (id, author, genre, publish_year, title) 
        VALUES (?, ?, ?, ?, ?)`,
		book.ID, book.Author, book.Genre, book.PublishYear, book.Title).Exec()

	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	iter := session.Query(`SELECT id, author, genre, publish_year, title FROM books`).Iter()

	var book models.Book
	for iter.Scan(&book.ID, &book.Author, &book.Genre, &book.PublishYear, &book.Title) {
		books = append(books, book)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user.ID = gocql.TimeUUID()
	err := session.Query(`
        INSERT INTO users (id, email, name) 
        VALUES (?, ?, ?)`,
		user.ID, user.Email, user.Name).Exec()

	if err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	iter := session.Query(`SELECT id, email, name FROM users`).Iter()

	var user models.User
	for iter.Scan(&user.ID, &user.Email, &user.Name) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	var borrow models.Borrow
	if err := json.NewDecoder(r.Body).Decode(&borrow); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := session.Query(`INSERT INTO borrowed_books (book_id, user_id, borrow_date) 
        VALUES (?, ?, ?)`, borrow.BookID, borrow.UserID, borrow.BorrowDate).Exec(); err != nil {
		http.Error(w, "Failed to borrow book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(borrow)
}

func GetBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	var borrowed []models.Borrow
	iter := session.Query(`SELECT book_id, user_id, borrow_date FROM borrowed_books`).Iter()

	var borrow models.Borrow
	for iter.Scan(&borrow.BookID, &borrow.UserID, &borrow.BorrowDate) {
		borrowed = append(borrowed, borrow)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, "Failed to fetch borrowed books", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrowed)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	var borrow models.Borrow

	if err := json.NewDecoder(r.Body).Decode(&borrow); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := session.Query(`DELETE FROM borrowed_books WHERE book_id = ? AND user_id = ?`,
		borrow.BookID, borrow.UserID).Exec()

	if err != nil {
		http.Error(w, "Failed to return book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("id")

	err := session.Query(`DELETE FROM books WHERE id = ?`, bookID).Exec()
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	err := session.Query(`DELETE FROM users WHERE id = ?`, userID).Exec()
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

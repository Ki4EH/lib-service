package repository

import (
	"database/sql"
	"github.com/Ki4EH/lib-service/catalog/entities"
	"log"
)

func InsertAuthor(db *sql.DB, name string) int {
	var id int
	err := db.QueryRow("INSERT INTO author (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func InsertBook(db *sql.DB, book entities.Book, authorId int) int {
	var id int
	err := db.QueryRow("INSERT INTO book (name, author_id, \"ISBN\") VALUES ($1, $2, $3) RETURNING id",
		book.Title, authorId, book.ISBN,
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func InsertCatalog(db *sql.DB, book entities.Book) {
	_, err := db.Exec("INSERT INTO catalog (book_id, count) VALUES ($1, $2)",
		book.ID, book.Count)
	if err != nil {
		panic(err)
	}
}

func UpdateCatalog(db *sql.DB, book entities.Book, count int) {
	_, err := db.Exec("UPDATE catalog SET count = $1 WHERE book_id = $2",
		book.Count+count, book.ID)
	if err != nil {
		panic(err)
	}
}
func DeleteFromGenreBook(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM genre_book WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table genre_book where book_id = %d", bookID)
}

func DeleteFromCatalog(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM catalog WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table catalog where book_id = %d", bookID)
}

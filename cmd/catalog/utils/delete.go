package utils

import (
	"database/sql"
	"log"
)

func DeleteBook(db *sql.DB, id int) {
	if !bookExists(db, id) {
		log.Printf("Book not found, nothing to delete")
		return
	}

	deleteFromGenreBook(db, id)
	deleteFromCatalog(db, id)

	_, err := db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return
	}
	log.Printf("Book deleted from database")
}

func deleteFromGenreBook(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM genre_book WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table genre_book where book_id = %d", bookID)
}

func deleteFromCatalog(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM catalog WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table catalog where book_id = %d", bookID)
}

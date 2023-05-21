package handler

import (
	"database/sql"
	"github.com/Ki4EH/lib-service/catalog/internal/repository"
	"log"
)

func DeleteBook(db *sql.DB, id int) {
	if !bookExists(db, id) {
		log.Printf("Book not found, nothing to delete")
		return
	}

	repository.DeleteFromGenreBook(db, id)
	repository.DeleteFromCatalog(db, id)

	_, err := db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return
	}
	log.Printf("Book deleted from database")
}

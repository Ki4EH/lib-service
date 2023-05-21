package handler

import (
	"database/sql"
	"fmt"
	"github.com/Ki4EH/lib-service/catalog/entities"
	"github.com/Ki4EH/lib-service/catalog/internal/repository"
)

func PostBook(db *sql.DB, book entities.Book) {
	b := findByISBN(db, book.ISBN)
	if b.ID != 0 {
		fmt.Println("Книга уже есть в каталоге!")
		book.ID = b.ID
		repository.UpdateCatalog(db, book, b.Count)
		return
	}

	authorId := getAuthorId(db, book.Author)
	if authorId == 0 {
		authorId = repository.InsertAuthor(db, book.Author)
	}

	id := repository.InsertBook(db, book, authorId)
	book.ID = id
	repository.InsertCatalog(db, book)
}

package utils

import (
	"database/sql"
	"fmt"
)

func PostBook(db *sql.DB, book Book) {
	b := findByISBN(db, book.ISBN)
	if b.ID != 0 {
		fmt.Println("Книга уже есть в каталоге!")
		book.ID = b.ID
		updateCatalog(db, book, b.Count)
		return
	}

	authorId := getAuthorId(db, book.Author)
	if authorId == 0 {
		authorId = insertAuthor(db, book.Author)
	}

	id := insertBook(db, book, authorId)
	book.ID = id
	insertCatalog(db, book)
}

func insertAuthor(db *sql.DB, name string) int {
	var id int
	err := db.QueryRow("INSERT INTO author (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func insertBook(db *sql.DB, book Book, authorId int) int {
	var id int
	err := db.QueryRow("INSERT INTO book (name, author_id, \"ISBN\") VALUES ($1, $2, $3) RETURNING id", book.Title, authorId, book.ISBN).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func updateCatalog(db *sql.DB, book Book, count int) {
	_, err := db.Exec("UPDATE catalog SET count = $1 WHERE book_id = $2", book.Count+count, book.ID)
	if err != nil {
		panic(err)
	}
}

func insertCatalog(db *sql.DB, book Book) {
	_, err := db.Exec("INSERT INTO catalog (book_id, count) VALUES ($1, $2)", book.ID, book.Count)
	if err != nil {
		panic(err)
	}
}

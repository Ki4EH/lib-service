package utils

import (
	"database/sql"
	"log"
)

func bookExists(db *sql.DB, id int) bool {
	row := db.QueryRow("SELECT id FROM book WHERE id = $1", id)

	var foundId int
	err := row.Scan(&foundId)
	if err != nil {
		panic(err)
	}

	if foundId > 0 {
		log.Printf("Book with id %d exists", id)
		return true
	}
	return false
}

func scanGenres(db *sql.DB, book *Book) {
	// Получения жанров книги
	genres, err1 := db.Query("SELECT name FROM genre_book JOIN genre g ON g.id = genre_book.genre_id WHERE book_id = $1", book.ID)
	if err1 != nil {
		panic(err1)
	}
	// Занесение жанров в переменную book
	var gens []string
	for genres.Next() {
		var g string
		err := genres.Scan(&g)
		if err != nil {
			panic(err)
		}
		gens = append(gens, g)
	}
	book.Genres = gens
}

func GetBookByID(db *sql.DB, id int) Book {
	// Получение от бд основной инфы про книгу, заносится в переменную book
	rows := db.QueryRow("select b.id, b.name, b.\"ISBN\", a.name, count from catalog JOIN book b ON catalog.book_id = b.id JOIN author a on a.id = b.author_id where b.id = $1", id)
	book := Book{}
	err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)

	if err != nil {
		panic(err)
	}
	scanGenres(db, &book)

	return book
}

func findByISBN(db *sql.DB, ISBN string) Book {
	row := db.QueryRow("SELECT b.id, b.name, a.name, \"ISBN\", count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a on a.id = b.author_id WHERE \"ISBN\" = $1", ISBN)
	book := Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Count)
	if err != nil {
		panic(err)
	}
	return book
}

func getAuthorId(db *sql.DB, name string) int {
	var id int
	err := db.QueryRow("SELECT id FROM author WHERE name = $1", name).Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

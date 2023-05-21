package handler

import (
	"database/sql"
	"fmt"
	"github.com/Ki4EH/lib-service/catalog/entities"
	"log"
)

var cosCount = 5
var minCos = 0.99

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

func scanGenres(db *sql.DB, book *entities.Book) error {
	// Getting the book's genres
	genres, err := db.Query("SELECT name FROM genre_book JOIN genre g ON g.id = genre_book.genre_id WHERE book_id = $1", book.ID)
	if err != nil {
		return fmt.Errorf("error querying database: %v", err)
	}
	defer genres.Close()

	var gens []string
	for genres.Next() {
		var g string
		err := genres.Scan(&g)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		gens = append(gens, g)
	}

	if err = genres.Err(); err != nil {
		return fmt.Errorf("error iterating over rows: %v", err)
	}

	book.Genres = gens
	return nil
}

func GetBookByID(db *sql.DB, id int) (entities.Book, error) {
	// Получение от бд основной инфы про книгу, заносится в переменную book
	rows := db.QueryRow("select b.id, b.name, b.\"ISBN\", a.name, count from catalog JOIN book b ON catalog.book_id = b.id JOIN author a on a.id = b.author_id where b.id = $1", id)
	book := entities.Book{}
	err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle empty response (no book found)
			return entities.Book{}, fmt.Errorf("book not found")
		}
		// Handle other errors
		return entities.Book{}, err
	}

	scanGenres(db, &book)

	return book, nil
}

func findByISBN(db *sql.DB, ISBN string) entities.Book {
	row := db.QueryRow("SELECT b.id, b.name, a.name, \"ISBN\", count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a on a.id = b.author_id WHERE \"ISBN\" = $1", ISBN)
	book := entities.Book{}
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
func contains(books []entities.Book, b entities.Book) bool {
	for _, a := range books {
		if a.ID == b.ID {
			return true
		}
	}
	return false
}

func getBooksFromCosMap(m map[float64][]entities.Book) []entities.Book {
	var books []entities.Book
	for len(books) < cosCount {
		if len(m) != 0 {
			cosMax := maxEl(m)
			for j := 0; j < len(m[cosMax]); j++ {
				if len(books) < cosCount {
					if contains(books, m[cosMax][j]) {
						continue
					}
					books = append(books, m[cosMax][j])
				}
			}
			delete(m, cosMax)
		}
		if len(books) >= cosCount {
			break
		}
	}

	return books
}

package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ki4EH/lib-service/catalog/entities"
	"strings"
)

// Search Через эту функцию будет осуществляться основной поиск в бд.
// Внутри нее будут вызываться остальные функции, связанные с поиском,
// со временем ее функционал будет наращиваться.
func Search(db *sql.DB, title string, author string) ([]entities.Book, error) {
	var result []entities.Book

	if title != "" {
		if author == "" {
			// Translate all elements from the search result by name into the result array
			var err error
			result, err = searchByTitle(db, title)
			if err != nil {
				return nil, err
			}
		} else {
			byTitle, err := searchByTitle(db, title)
			if err != nil {
				return nil, err
			}

			byAuthor, err := searchByAuthor(db, author)
			if err != nil {
				return nil, err
			}

			for _, b := range byTitle {
				if contains(result, b) {
					continue
				}
				result = append(result, b)
			}

			for _, b := range byAuthor {
				if contains(result, b) {
					continue
				}
				result = append(result, b)
			}
		}
	} else {
		if author != "" {
			var err error
			result, err = searchByAuthor(db, author)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("both title and author parameters are empty")
		}
	}

	for i := range result {
		b := &result[i]
		err := scanGenres(db, b)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Поиск по названию :\
func searchByTitle(db *sql.DB, title string) ([]entities.Book, error) {
	var books []entities.Book
	m := make(map[float64][]entities.Book)
	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON a.id = b.author_id")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.Book

		err = rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		words := strings.Split(strings.ToLower(book.Title), " ")
		qWords := strings.Split(strings.ToLower(title), " ")
		for _, w := range words {
			for _, qw := range qWords {
				cosValue := cosR(w, qw)
				if cosValue >= minCos {
					if m[cosValue] == nil {
						var bo []entities.Book
						bo = append(bo, book)
						m[cosValue] = bo
					} else {
						m[cosValue] = append(m[cosValue], book)
					}
				}
			}
		}

	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	books = getBooksFromCosMap(m)
	return books, nil
}

func searchByAuthor(db *sql.DB, author string) ([]entities.Book, error) {
	var books []entities.Book
	var m = make(map[float64][]entities.Book)

	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON b.author_id = a.id")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.Book

		err = rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		words := strings.Split(strings.ToLower(book.Author), " ")
		qWords := strings.Split(strings.ToLower(author), " ")
		for _, w := range words {
			for _, qw := range qWords {
				cosValue := cosR(w, qw)
				if cosValue >= minCos {
					if m[cosValue] == nil {
						var bo []entities.Book
						bo = append(bo, book)
						m[cosValue] = bo
					} else {
						m[cosValue] = append(m[cosValue], book)
					}
				}
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	books = getBooksFromCosMap(m)
	return books, nil
}

package utils

import (
	"database/sql"
	"strings"
)

// Search Через эту функцию будет осуществляться основной поиск в бд.
// Внутри нее будут вызываться остальные функции, связанные с поиском,
// со временем ее функционал будет наращиваться.
func Search(db *sql.DB, title string, author string) []Book {
	var result []Book

	if title != "" {
		if author == "" {
			// Трансляция всех элементов из резяльтата поиска по названию в массив result
			result = searchByTitle(db, title)
		}
	} else {
		if author != "" {
			result = searchByAuthor(db, author)
		} else {
			result = nil
		}
	}

	// Здесь будут применяться другие функции поиска,
	// которые будут вносить изменения в массив result.
	// Возможно эту систему поиска потом поменяем

	for i := range result {
		b := &result[i]
		scanGenres(db, b)
	}

	return result
}

// Поиск по названию :\
func searchByTitle(db *sql.DB, title string) []Book {
	var books []Book
	var m map[float64][]Book
	m = make(map[float64][]Book)
	minCos := 0.9
	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON a.id = b.author_id")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			panic(err)
		}

		cosValue := cosR(strings.ToLower(book.Title), title)
		if cosValue >= minCos {
			if m[cosValue] == nil {
				var bo []Book
				bo = append(bo, book)
				m[cosValue] = bo
			} else {
				m[cosValue] = append(m[cosValue], book)
			}

		}

	}
	for i := 0; i < 5; i++ {
		if len(m) != 0 {
			cosMax := maxEl(m)
			for j := 0; j < len(m[cosMax]); j++ {
				if len(books) < 5 {
					books = append(books, m[cosMax][j])
				}
			}
			delete(m, cosMax)

		}
		if len(books) >= 5 {
			break
		}
	}
	return books
}

func searchByAuthor(db *sql.DB, author string) []Book {
	var books []Book
	var m = make(map[float64][]Book)
	minCos := 0.9
	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON b.author_id = a.id")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			panic(err)
		}

		cosValue := cosR(strings.ToLower(book.Author), author)

		if cosValue >= minCos {
			if m[cosValue] == nil {
				var bo []Book
				bo = append(bo, book)
				m[cosValue] = bo
			} else {
				m[cosValue] = append(m[cosValue], book)
			}
		}

	}
	for i := 0; i < 5; i++ {
		if len(m) != 0 {
			cosMax := maxEl(m)
			for j := 0; j < len(m[cosMax]); j++ {
				if len(books) < 5 {
					books = append(books, m[cosMax][j])
				}
			}
			delete(m, cosMax)

		}
		if len(books) >= 5 {
			break
		}
	}
	return books
}

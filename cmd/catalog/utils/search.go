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

		words := strings.Split(strings.ToLower(book.Title), " ")
		qWords := strings.Split(strings.ToLower(title), " ")
		for _, w := range words {
			for _, qw := range qWords {
				cosValue := cosR(w, qw)
				println(w, qw, cosValue)
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
		}

	}
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

func searchByAuthor(db *sql.DB, author string) []Book {
	var books []Book
	var m = make(map[float64][]Book)
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

		words := strings.Split(strings.ToLower(book.Author), " ")
		qWords := strings.Split(strings.ToLower(author), " ")
		for _, w := range words {
			for _, qw := range qWords {
				cosValue := cosR(w, qw)
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
		}

	}
	for i := 0; i < cosCount; i++ {
		if len(m) != 0 {
			cosMax := maxEl(m)
			for j := 0; j < len(m[cosMax]); j++ {
				if len(books) < cosCount {
					books = append(books, m[cosMax][j])
				}
			}
			delete(m, cosMax)
		}
	}
	return books
}

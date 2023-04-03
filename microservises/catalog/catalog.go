package catalog

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

// RunCatalogHandler Описывает обработку http запросов
func RunCatalogHandler(db *sql.DB) {
	http.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(request.URL.Query().Get("book_id"))
		if err != nil {
			panic(err)
		}
		book := getBookByID(db, id)

		res, err1 := json.Marshal(book)
		if err1 != nil {
			panic(err1)
		}
		_, err2 := writer.Write(res)
		if err2 != nil {
			return
		}
	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("title")
		author := request.URL.Query().Get("author")
		books := search(db, name, author)

		js, err := json.Marshal(books)
		if err != nil {
			panic(err)
		}
		_, err1 := writer.Write(js)
		if err1 != nil {
			return
		}
	})
	//http.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
	//	if request.Method == "POST" {
	//	}
	//})
}

func getBookByID(db *sql.DB, id int) Book {
	// Получение от бд основной инфы про книгу, заносится в переменную book
	rows := db.QueryRow("select * from books where id = $1", id)
	book := Book{}
	var nullAuthor sql.NullString
	err := rows.Scan(&book.ID, &book.Title, &book.Count, &nullAuthor)
	if reflect.TypeOf(nullAuthor) == nil {
		book.Author = ""
	} else {
		book.Author = nullAuthor.String
	}

	if err != nil {
		panic(err)
	}
	scanGenres(db, &book)

	return book
}

// Через эту функцию будет осуществляться основной поиск в бд.
// Внутри нее будут вызываться остальные функции, связанные с поиском,
// со временем ее функционал будет наращиваться.
func search(db *sql.DB, title string, author string) []Book {
	var result []Book

	if title != "" {
		if author == "" {
			// Трансляция всех элементов из резяльтата поиска по названию в массив result
			result = searchByTitle(db, title)
		}
		//else {
		//	result = searchByTnA(db, title, author)
		//}
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
	rows, err := db.Query("SELECT books.id, title, count, name FROM books JOIN authors a ON a.id = books.author_id WHERE title = $1", title)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Count, &book.Author)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func searchByAuthor(db *sql.DB, author string) []Book {
	var books []Book
	rows, err := db.Query("SELECT books.id, title, count, name FROM books JOIN authors a ON books.author_id = a.id WHERE a.name = $1", author)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Count, &book.Author)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func scanGenres(db *sql.DB, book *Book) {
	// Получения жанров книги
	genres, err1 := db.Query("SELECT name FROM genre_book JOIN genres g ON g.id = genre_book.genre_id WHERE book_id = $1", book.ID)
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

//
//func getAllBooks(db *sql.DB) []Book {
//	var res []Book
//
//	rows, err := db.Query("SELECT * FROM books")
//	if err != nil {
//		panic(err)
//	}
//
//	for rows.Next() {
//		var b = Book{}
//		rows.Scan(&b.ID, &b.Title, &b.Count)
//		res = append(res, b)
//	}
//
//	rows, err1 := db.Query("SELECT book_id, name FROM genre_book JOIN genres g ON genre_book.genre_id = g.id WHERE book_id = 1;")
//	if err1 != nil {
//		panic(err1)
//	}
//	for rows.Next() {
//		var id int
//		var genre string
//		rows.Scan(&id, &genre)
//
//		for i := range res {
//			if res[i].ID == id {
//				res[i].Genres = append(res[i].Genres, genre)
//			}
//		}
//	}
//	return res
//}

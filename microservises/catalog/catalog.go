package catalog

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int      `json:"id"`
	Author string   `json:"authors"`
	Title  string   `json:"title"`
	ISBN   string   `json:"isbn"`
	Count  int      `json:"count"`
	Genres []string `json:"genres"`
}

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
		writer.Write(res)
	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")
		books := search(db, name)

		js, err := json.Marshal(books)
		if err != nil {
			panic(err)
		}
		writer.Write(js)
	})
}

func getBookByID(db *sql.DB, id int) Book {
	// Получение от бд основной инфы про книгу, заносится в переменную book
	rows := db.QueryRow("select * from books where id = $1", id)
	book := Book{}
	err := rows.Scan(&book.ID, &book.Title, &book.Count)
	if err != nil {
		panic(err)
	}

	// Получения жанров книги
	genres, err1 := db.Query("SELECT name FROM genre_book JOIN genres g ON g.id = genre_book.genre_id WHERE book_id = $1", id)
	if err1 != nil {
		panic(err1)
	}

	// Занесение жанров в переменную book
	var gens []string
	for genres.Next() {
		var g string
		genres.Scan(&g)
		gens = append(gens, g)
	}
	book.Genres = gens

	return book
}

// Через эту функцию будет осуществляться основной поиск в бд.
// Внутри нее будут вызываться остальные функции, связанные с поиском,
// со временем ее функционал будет наращиваться.
func search(db *sql.DB, title string) []Book {
	var result []Book

	// Трансляция всех элементов из резяльтата поиска по названию в массив result
	for _, elem := range searchByTitle(db, title) {
		result = append(result, elem)
	}

	// Здесь буут применяться другие функции поиска,
	// которые будут вносить изменения в массив result.
	// Возможно эту систему поиска потом поменяем

	return result
}

// Поиск по названию :\
func searchByTitle(db *sql.DB, title string) []Book {
	var books []Book
	rows, err := db.Query("SELECT * FROM books WHERE title = $1", title)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.Count)
		books = append(books, book)
	}

	return books
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

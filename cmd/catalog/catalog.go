package catalog

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// RunCatalogHandler Описывает обработку http запросов
func RunCatalogHandler(db *sql.DB) {
	http.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
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
		} else if request.Method == http.MethodPost {
			var book Book = Book{
				Title:  request.URL.Query().Get("title"),
				Author: request.URL.Query().Get("author"),
				ISBN:   request.URL.Query().Get("isbn")}
			c, err := strconv.Atoi(request.URL.Query().Get("count"))
			if err != nil {
				panic(err)
			}
			book.Count = c
			postBook(db, book)
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
	rows := db.QueryRow("select b.id, b.name, b.\"ISBN\", a.name, count from catalog JOIN book b ON catalog.book_id = b.id JOIN author a on a.id = b.author_id where b.id = $1", id)
	book := Book{}
	err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
	//if reflect.TypeOf(nullAuthor) == nil {
	//	book.Author = ""
	//} else {
	//	book.Author = nullAuthor.String
	//}

	if err != nil {
		panic(err)
	}
	scanGenres(db, &book)

	return book
}

func postBook(db *sql.DB, book Book) {
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
	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON a.id = b.author_id WHERE b.name = $1", title)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func searchByAuthor(db *sql.DB, author string) []Book {
	var books []Book
	rows, err := db.Query("SELECT b.id, b.name, \"ISBN\", a.name, count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a ON b.author_id = a.id WHERE a.name = $1", author)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Count)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func findByISBN(db *sql.DB, ISBN string) Book {
	row := db.QueryRow("SELECT b.id, b.name, a.name, \"ISBN\", count FROM catalog JOIN book b on b.id = catalog.book_id JOIN author a on a.id = b.author_id WHERE \"ISBN\" = $1", ISBN)
	book := Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Count)
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

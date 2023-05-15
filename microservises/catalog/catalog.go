package catalog

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// RunCatalogHandler Описывает обработку http запросов
func RunCatalogHandler(db *sql.DB) {
	http.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if request.Method == http.MethodGet {
			id, err := strconv.Atoi(request.URL.Query().Get("book_id"))
			if err != nil {
				http.Error(writer, http.StatusText(400), 400)
				panic(err)
				return
			}
			book := getBookByID(db, id)

			if book.ID == 0 {
				http.Error(writer, http.StatusText(404), 404)
			} else {
				res, err1 := json.Marshal(book)
				if err1 != nil {
					http.Error(writer, http.StatusText(500), 500)
					panic(err1)
				}
				_, err2 := writer.Write(res)
				if err2 != nil {
					http.Error(writer, http.StatusText(500), 500)
					panic(err2)
				}
			}
		} else if request.Method == http.MethodPost {
			//user identity
			header := request.Header.Get("Authorization")
			if header == "" {
				http.Error(writer, http.StatusText(401), 401)
				panic("NO AUTHORIZATION TOKEN")
			}
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 {
				http.Error(writer, http.StatusText(401), 401)
				panic("invalid token")
			}
			_, role, err := ParseToken(headerParts[1])
			if err != nil {
				panic(err)
			}
			if role != 1 {
				http.Error(writer, http.StatusText(403), 403)
				panic("Forbidden")
			}
			var book Book = Book{
				Title:  request.URL.Query().Get("title"),
				Author: request.URL.Query().Get("author"),
				ISBN:   request.URL.Query().Get("isbn")}
			book.Genres = strings.Split(request.URL.Query().Get("genres"), ",")
			c, err := strconv.Atoi(request.URL.Query().Get("count"))
			if err != nil {
				http.Error(writer, http.StatusText(400), 400)
				panic(err)
				return
			} else {
				book.Count = c
				postBook(db, book)
			}
		} else if request.Method == http.MethodDelete {
			id, err := strconv.Atoi(request.URL.Query().Get("id"))
			if err != nil {
				http.Error(writer, http.StatusText(400), 400)
				panic(err)
			}
			deleteBook(db, id)
		}
	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("title")
		author := request.URL.Query().Get("author")
		query := request.URL.Query().Get("query")

		if name == "" && author == "" {
			var titlePretender, authorPretender []string
			handleQuery(query)

			for _, str := range titlePretender {
				_, err1 := writer.Write([]byte(str))
				if err1 != nil {
					return
				}
			}
			_, err1 := writer.Write([]byte("\n"))
			if err1 != nil {
				return
			}
			for _, str := range authorPretender {
				_, err := writer.Write([]byte(str))
				if err != nil {
					return
				}
			}
			return
		}

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
}

func getBookByID(db *sql.DB, id int) Book {
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

//
//func handleQuery(query string) ([]string, []string) {
//	var titlePretender, authorPretender []string
//	doc, err := prose.NewDocument(query)
//	if err != nil {
//		panic(err)
//	}
//
//	for _, token := range doc.Tokens() {
//		titlePretender = append(titlePretender, token.Text)
//
//		if token.Label == "B-PERSON" {
//			authorPretender = append(authorPretender, token.Text)
//		}
//	}
//
//	for _, str := range titlePretender {
//		print(str + ", ")
//	}
//	println()
//	for _, str := range authorPretender {
//		print(str + ", ")
//	}
//
//	tempTitles := combinations.All(titlePretender)
//	tempAuthors := combinations.All(authorPretender)
//	titlePretender = []string{}
//	authorPretender = []string{}
//
//	for _, i := range tempTitles {
//		var str = ""
//		for _, j := range i {
//			str +=
//		}
//	}
//	return tempTitles, tempAuthors
//
//	return titlePretender, authorPretender
//}

func deleteBook(db *sql.DB, id int) {
	if !bookExists(db, id) {
		log.Printf("Book not found, nothing to delete")
		return
	}

	deleteFromGenreBook(db, id)
	deleteFromCatalog(db, id)

	_, err := db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return
	}
	log.Printf("Book deleted from database")
}

func deleteFromGenreBook(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM genre_book WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table genre_book where book_id = %d", bookID)
}

func deleteFromCatalog(db *sql.DB, bookID int) {
	_, err := db.Exec("DELETE FROM catalog WHERE book_id = $1", bookID)
	if err != nil {
		return
	}
	log.Printf("Row deleted from table catalog where book_id = %d", bookID)
}

func bookExists(db *sql.DB, id int) bool {
	row := db.QueryRow("SELECT id FROM book WHERE id = $1", id)

	var foundId int
	row.Scan(&foundId)

	if foundId > 0 {
		log.Printf("Book with id %d exists", id)
		return true
	}
	return false
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

		cosValue := cos_r(book.Title, title)
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
			cosMax := max_el(m)
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
	var m map[float64][]Book = make(map[float64][]Book)
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

		cosValue := cos_r(book.Author, author)

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
			cosMax := max_el(m)
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

func cos_r(s1 string, s2 string) float64 {

	var sum1, sum2, sum3 int
	n := int(math.Min(float64(len(s1)), float64(len(s2))))
	for i := 0; i < n; i++ {
		sum1 += (int(rune(s1[i])) + 200) * (200 + int(rune(s2[i])))
		sum2 += (int(rune(s1[i])) + 200) * (200 + int(rune(s1[i])))
		sum3 += (200 + int(rune(s2[i]))) * (200 + int(rune(s2[i])))
		//fmt.Println(rune(s1[i]), rune(s2[i]))
	}
	//fmt.Println(sum1, sum2, sum3)
	cos := float64(sum1) / (math.Sqrt(float64(sum2)) * math.Sqrt(float64(sum3)))
	//fmt.Println(s1, s2, cos)
	return cos
}

func max_el(m map[float64][]Book) float64 {
	var i float64
	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	i = keys[len(m)-1]
	return i
}

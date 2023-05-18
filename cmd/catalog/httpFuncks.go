package catalog

import (
	"Catalog/cmd/catalog/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func httpBookGet(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	id, err := strconv.Atoi(request.URL.Query().Get("book_id"))
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)
		panic(err)
		return
	}
	book := utils.GetBookByID(db, id)

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
}
func httpBookPost(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
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
	var book = utils.Book{
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
		utils.PostBook(db, book)
	}
}

func httpBookDelete(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)
		panic(err)
	}
	utils.DeleteBook(db, id)
}

func httpSearch(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	name := strings.ToLower(request.URL.Query().Get("title"))
	author := strings.ToLower(request.URL.Query().Get("author"))

	if name == "" && author == "" {
		var titlePretender, authorPretender []string
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

	books := utils.Search(db, name, author)

	js, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	_, err1 := writer.Write(js)
	if err1 != nil {
		return
	}
}

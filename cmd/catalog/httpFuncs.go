package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Ki4EH/lib-service/catalog/entities"
	"github.com/Ki4EH/lib-service/catalog/handler"
	"net/http"
	"strconv"
	"strings"
)

func httpBookGet(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	// Convert the id to an integer
	id, err := strconv.Atoi(request.URL.Query().Get("book_id"))
	if err != nil {
		http.Error(writer, "Invalid book_id parameter", http.StatusBadRequest)
		return
	}

	// Fetch the book from the database
	book, err := handler.GetBookByID(db, id)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(writer, "Book not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Convert the book to JSON
	jsonBook, err := json.Marshal(book)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	writer.Header().Set("Content-Type", "application/json")

	// Write the JSON book to the response
	_, err = writer.Write(jsonBook)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
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
	var book = entities.Book{
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
		handler.PostBook(db, book)
	}
}

func httpBookDelete(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)
		panic(err)
	}
	handler.DeleteBook(db, id)
}

func httpSearch(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	name := strings.ToLower(request.URL.Query().Get("title"))
	author := strings.ToLower(request.URL.Query().Get("author"))

	if name == "" && author == "" {
		// This is where you should handle the case where both 'name' and 'author' are empty
		http.Error(writer, "Missing title and author parameters", http.StatusBadRequest)
		return
	}

	books, err := handler.Search(db, name, author)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonBooks, err := json.Marshal(books)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(jsonBooks)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

package main

import (
	"database/sql"
	"net/http"
)

// RunCatalogHandler Описывает обработку http запросов
func RunCatalogHandler(db *sql.DB) {
	http.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,GET,UPDATE")
		if request.Method == http.MethodGet {
			httpBookGet(writer, request, db)
		} else if request.Method == http.MethodPost {
			httpBookPost(writer, request, db)
		} else if request.Method == http.MethodDelete {
			httpBookDelete(writer, request, db)
		}
	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		httpSearch(writer, request, db)
	})
}

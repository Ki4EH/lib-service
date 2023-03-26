package main

import (
	"Catalog/microservises/catalog"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	// Подключение у бд
	connStr := "user=postgres password=livmas dbname=testcatalog sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	catalog.RunCatalogHandler(db)
	fmt.Println(http.ListenAndServe(":5000", nil))
}

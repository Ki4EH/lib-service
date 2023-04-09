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
	connStr := "host=95.140.159.168 port=5433 user=go_project password=rIo3Fc dbname=lib-service-test sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	catalog.RunCatalogHandler(db)
	fmt.Println(http.ListenAndServe(":5000", nil))
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// Retrieve secrets from environment variables
	host := "95.140.159.168"
	port := "5433"
	user := "go_project"
	password := "rIo3Fc"
	dbname := "lib-service-test"

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Println("Connecting to PostgreSQL database")
	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	RunCatalogHandler(db)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

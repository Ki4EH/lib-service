package catalog

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/example/catalog"
	_ "github.com/lib/pq"
)

func main() {
	// Retrieve secrets from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	catalog.RunCatalogHandler(db)

	fmt.Println(http.ListenAndServe("127.0.0.1:8080", nil))
}

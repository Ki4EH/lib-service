package handlers

import (
	"database/sql"
	"fmt"
	"log"
	
	"github.com/Ki4EH/lib-service/stat-n-rec/entities"
	
	_ "github.com/lib/pq"
)

func GetAuthorsStats() {
	// Подключение к БД PostgreSQL
	connStr := "user=your_user password=your_password dbname=your_database host=your_host port=your_port sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получение списка прочитанных книг из БД
	rows, err := db.Query("SELECT id, title, read FROM books WHERE read=true")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var readBooks []entities.Book_t
	for rows.Next() {
		var book entities.Book_t
		if err := rows.Scan(&book.id, &book.title, &book.read); err != nil {
			log.Fatal(err)
		}
		readBooks = append(readBooks, book)
	}

	// Получение списка авторов
	rows, err = db.Query("SELECT id, name FROM authors")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var authors []entities.Author_t
	for rows.Next() {
		var author entities.Author_t
		if err := rows.Scan(&author.id, &author.name); err != nil {
			log.Fatal(err)
		}
		authors = append(authors, author)
	}

	// Создание словаря с количеством прочитанных книг по авторам
	authorsStats := make(map[string]int)
	for _, author := range authors {
		numReadBooks := 0
		for _, book := range readBooks {
			row := db.QueryRow("SELECT author_id FROM books WHERE id=$1", book.id)
			var bookAuthorId int
			if err := row.Scan(&bookAuthorId); err != nil {
				log.Fatal(err)
			}
			if bookAuthorId == author.id {
				numReadBooks++
			}
		}
		authorsStats[author.name] = numReadBooks
	}

	// Вывод статистики по авторам
	for author, numReadBooks := range authorsStats {
		fmt.Printf("Автор %s - количество прочитанных книг: %d\n", author, numReadBooks)
	}
}

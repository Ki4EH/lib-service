package entities

// Структура для передачи данных
type json_t struct {
	message string
	details []int
}

// Структура Книги для статистики
type Book_t struct {
	id    int
	title string
	read  bool
}

// Структура Автора для статистики
type Author_t struct {
	id   int
	name string
}

package utils

type Book struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	ISBN   string   `json:"isbn"`
	Count  int      `json:"count"`
	Genres []string `json:"genres"`
}

func contains(books []Book, b Book) bool {
	for _, a := range books {
		if a.ID == b.ID {
			return true
		}
	}
	return false
}

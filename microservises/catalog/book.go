package catalog

type Book struct {
	ID     int      `json:"id"`
	Author string   `json:"author"`
	Title  string   `json:"title"`
	ISBN   string   `json:"isbn"`
	Count  int      `json:"count"`
	Genres []string `json:"genres"`
}

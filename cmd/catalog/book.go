package main

type Book struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	ISBN   string   `json:"isbn"`
	Count  int      `json:"count"`
	Genres []string `json:"genres"`
}

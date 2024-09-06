package models

type Author_Book struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	BookID     int    `json:"book_id"`
	BookName   string `json:"book_name"`
}

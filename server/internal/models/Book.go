package model

type Book struct {
	Title    string `json:"title"`
	ISBN     string `json:"isbn"`
	AuthorID int    `json:"author_id"`
	Price    int    `json:"price"`
}

func NewBook(title string, isbn string, authorID int, price int) *Book {
	return &Book{
		Title:    title,
		ISBN:     isbn,
		AuthorID: authorID,
		Price:    price,
	}
}

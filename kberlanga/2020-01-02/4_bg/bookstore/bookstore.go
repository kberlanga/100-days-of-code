package bookstore

// Book represents information about a book.
type Book struct {
	Title           string
	Authors         []string
	Description     string
	Copies          int
	Serie           int
	PriceCents      int
	DiscountPercent float64
	Featured        bool
}

func GetAllBooks(catalog []Book) (allBooks []Book) {
	// var allBooks []Book
	for _, book := range catalog {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

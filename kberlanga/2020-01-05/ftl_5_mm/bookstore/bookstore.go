package bookstore

import (
	"strconv"

	"github.com/segmentio/ksuid"
)

// Book represents information about a book.
type Book struct {
	ID              string
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
	allBooks = make([]Book, 0)
	for _, book := range catalog {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

func AddBookToCatalog(newBook Book, catalog []Book) (newCatalog []Book) {
	newCatalog = append(catalog, newBook)
	return newCatalog
}

func NewID() (ID string) {
	ID = ksuid.New().String()
	return ID
}

func GetBookDetails(book Book) (details string) {
	details = "The book " + book.Title + " with ID " + book.ID + " has a cost of " + strconv.Itoa(book.PriceCents)
	return details
}

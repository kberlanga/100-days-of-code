package bookstore2

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

func GetAllBooks(catalog map[string]Book) (allBooks map[string]Book) {
	allBooks = map[string]Book{}
	for key, book := range catalog {
		allBooks[key] = book
	}
	return allBooks
}

func AddBookToCatalog(newBook Book, catalog map[string]Book) (newCatalog map[string]Book) {
	newCatalog = catalog
	// for key, book := range catalog {
	// 	newCatalog[key] = book
	// }
	ID := "Book" + strconv.Itoa(len(catalog)+1)
	newCatalog[ID] = newBook
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

func GetAllByAuthor(author string, catalog []Book) (allBooks []Book) {
	allBooks = make([]Book, 0)
	for _, book := range catalog {
		for _, myAuthor := range book.Authors {
			if myAuthor == author {
				allBooks = append(allBooks, book)
			}
		}
	}
	return allBooks
}

func GetCatalogDetails(catalog []Book) (allDetails string) {
	allDetails = ""
	for i, book := range catalog {
		lon := len(catalog) - 1
		if lon != i {
			allDetails += GetBookDetails(book) + "\n"
		} else {
			allDetails += GetBookDetails(book)
		}
	}
	return allDetails
}

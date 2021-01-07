package bookstore2_test

import (
	"bookstore2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestBook(t *testing.T) {
// 	_ = map[string]bookstore2.Book{
// "Book1": {
// 	ID:              "abcde123",
// 		Title:           "Spark Joy",
// 		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
// 		Description:     "A tiny, cheerful Japanese woman explains tidying.",
// 		Copies:          30,
// 		Serie:           1,
// 		PriceCents:      1199,
// 		DiscountPercent: 0.05,
// 		Featured:        true,
// 	}
// 	},
// }

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Get all books of empty catalog", catalog: map[string]bookstore2.Book{}, want: map[string]bookstore2.Book{}},
		{name: "Get all books of catalog", catalog: books, want: books},
		{name: "Get all books of catalog with a element", catalog: book, want: book},
	}

	for _, tcase := range testCases {
		got := bookstore2.GetAllBooks(tcase.catalog)
		if !cmp.Equal(tcase.want, got) {
			t.Errorf("Case -> %s\nGetAllBooks(%v): \nwant %v, \ngot %v", tcase.name, tcase.catalog, tcase.want, got)
		}
	}
}

func TestAddBookToCatalog(t *testing.T) {
	t.Parallel()
	testCases := []testCaseAdd{
		{name: "Add new book to empty catalog", newBook: book["Book1"], catalog: map[string]bookstore2.Book{}, want: map[string]bookstore2.Book{"Book1": book["Book1"]}},
		{name: "Add new book to catalog with books", newBook: book["Book1"], catalog: books, want: newCatalog},
	}
	for _, tcase := range testCases {
		got := bookstore2.AddBookToCatalog(tcase.newBook, tcase.catalog)
		if !cmp.Equal(tcase.want, got) {
			t.Errorf("Case -> %s\nAddBookToCatalog(%v, %v): \nwant %v, \ngot %v", tcase.name, tcase.newBook, tcase.catalog, tcase.want, got)
		}
	}
}

// func TestGetBookDetails(t *testing.T) {
// 	t.Parallel()
// 	testCases := []testCaseDetails{
// 		{name: "Detail of book", book: book[0], want: "The book Spark Joy with ID abcde123 has a cost of 1199"},
// 		// {name: "Add new book to catalog with books", newBook: book[0], catalog: books, want: newCatalog},
// 	}
// 	for _, tcase := range testCases {
// 		got := bookstore2.GetBookDetails(tcase.book)
// 		if !cmp.Equal(tcase.want, got) {
// 			t.Errorf("Case -> %s\nGetBookDetails(%v): \nwant %v, \ngot %v", tcase.name, tcase.book, tcase.want, got)
// 		}
// 	}
// }

// func TestGetAllByAuthor(t *testing.T) {
// 	t.Parallel()
// 	testCases := []testCaseAuthor{
// 		{name: "only book by that author", author: "Karla Berlanga", catalog: books, want: booksByAuthor1},
// 		{name: "only book by some authors", author: "Marie Kondō", catalog: books, want: booksByAuthor2},
// 		{name: "only book by that author", author: "X", catalog: books2, want: booksByAuthor3},
// 		{name: "author not found", author: "My Author", catalog: books2, want: []bookstore2.Book{}},
// 	}
// 	for _, tcase := range testCases {
// 		got := bookstore2.GetAllByAuthor(tcase.author, tcase.catalog)
// 		if !cmp.Equal(tcase.want, got) {
// 			t.Errorf("Case -> %s\nGetBookDetails(%v, %v): \nwant %v, \ngot %v", tcase.name, tcase.author, tcase.catalog, tcase.want, got)
// 		}
// 	}
// }

// func TestGetCatalogDetails(t *testing.T) {
// 	t.Parallel()
// 	testCases := []testCaseString{
// 		{name: "catalog with three books", catalog: books2, want: case1},
// 		{name: "An empty catalog", catalog: []bookstore2.Book{}, want: case2},
// 		{name: "catalog with only book", catalog: book, want: case3},
// 	}
// 	for _, tcase := range testCases {
// 		got := bookstore2.GetCatalogDetails(tcase.catalog)
// 		if !cmp.Equal(tcase.want, got) {
// 			t.Errorf("Case -> %s\nGetCatalogDetails(%v): \nwant %v, \ngot %v", tcase.name, tcase.catalog, tcase.want, got)
// 		}
// 	}
// }

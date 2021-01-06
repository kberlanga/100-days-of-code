package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	name          string
	catalog, want []bookstore.Book
	expectedError bool
}

type testCaseAdd struct {
	name          string
	newBook       bookstore.Book
	catalog, want []bookstore.Book
	expectedError bool
}

type testCaseDetails struct {
	name          string
	book          bookstore.Book
	want          string
	expectedError bool
}

var books = []bookstore.Book{
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var book = []bookstore.Book{
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var newCatalog = []bookstore.Book{
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
	{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Get all books of empty catalog", catalog: []bookstore.Book{}, want: []bookstore.Book{}},
		{name: "Get all books of catalog", catalog: books, want: books},
		{name: "Get all books of catalog with a element", catalog: book, want: book},
	}

	for _, tcase := range testCases {
		got := bookstore.GetAllBooks(tcase.catalog)
		if !cmp.Equal(tcase.want, got) {
			t.Errorf("Case -> %s\nGetAllBooks(%v): \nwant %v, \ngot %v", tcase.name, tcase.catalog, tcase.want, got)
		}
	}
}

func TestAddBookToCatalog(t *testing.T) {
	t.Parallel()
	testCases := []testCaseAdd{
		{name: "Add new book to empty catalog", newBook: book[0], catalog: []bookstore.Book{}, want: []bookstore.Book{book[0]}},
		{name: "Add new book to catalog with books", newBook: book[0], catalog: books, want: newCatalog},
	}
	for _, tcase := range testCases {
		got := bookstore.AddBookToCatalog(tcase.newBook, tcase.catalog)
		if !cmp.Equal(tcase.want, got) {
			t.Errorf("Case -> %s\nAddBookToCatalog(%v, %v): \nwant %v, \ngot %v", tcase.name, tcase.newBook, tcase.catalog, tcase.want, got)
		}
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()
	testCases := []testCaseDetails{
		{name: "Detail of book", book: book[0], want: "The book Spark Joy with ID abcde123 has a cost of 1199"},
		// {name: "Add new book to catalog with books", newBook: book[0], catalog: books, want: newCatalog},
	}
	for _, tcase := range testCases {
		got := bookstore.GetBookDetails(tcase.book)
		if !cmp.Equal(tcase.want, got) {
			t.Errorf("Case -> %s\nGetBookDetails(%v): \nwant %v, \ngot %v", tcase.name, tcase.book, tcase.want, got)
		}
	}
}

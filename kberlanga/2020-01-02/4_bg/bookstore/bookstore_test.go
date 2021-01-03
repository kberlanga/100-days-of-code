package bookstore_test

import (
	"bookstore"
	"reflect"
	"testing"
)

type testCase struct {
	name          string
	catalog, want []bookstore.Book
	expectedError bool
}

func TestBook(t *testing.T) {
	_ = bookstore.Book{
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
	books := []bookstore.Book{
		{
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

	book := []bookstore.Book{
		{
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
	testCases := []testCase{
		{name: "Get all books of empty catalog", catalog: []bookstore.Book{}, want: []bookstore.Book{}},
		{name: "Get all books of catalog", catalog: books, want: books},
		{name: "Get all books of catalog with a element", catalog: book, want: book},
	}

	for _, tcase := range testCases {
		got := bookstore.GetAllBooks(tcase.catalog)
		if !reflect.DeepEqual(tcase.want, got) {
			// if tcase.want != got {
			t.Errorf("Case -> %s\nGetAllBooks(%v): \nwant %v, \ngot %v", tcase.name, tcase.catalog, tcase.want, got)
		}
	}
}

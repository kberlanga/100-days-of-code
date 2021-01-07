package bookstore2_test

import "bookstore2"

type testCase struct {
	name          string
	catalog, want map[string]bookstore2.Book
	expectedError bool
}

type testCaseString struct {
	name          string
	catalog       map[string]bookstore2.Book
	want          string
	expectedError bool
}

type testCaseAdd struct {
	name          string
	newBook       bookstore2.Book
	catalog, want map[string]bookstore2.Book
	expectedError bool
}

type testCaseDetails struct {
	name          string
	book          bookstore2.Book
	want          string
	expectedError bool
}

type testCaseAuthor struct {
	name          string
	author        string
	want, catalog map[string]bookstore2.Book
	expectedError bool
}

var books = map[string]bookstore2.Book{
	"Book1": {
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
	"Book2": {
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var book = map[string]bookstore2.Book{
	"Book1": {
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

var newCatalog = map[string]bookstore2.Book{
	"Book1": {
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
	"Book2": {
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
	"Book3": {
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

var books2 = map[string]bookstore2.Book{
	"Book1": {
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
	"Book2": {
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
	"Book3": {
		ID:              "abcde345",
		Title:           "Go BootCamp",
		Authors:         []string{"Matt Aimonetti", "X", "Y", "Z"},
		Description:     "Everything you need to know to get started with Go",
		Copies:          40,
		Serie:           1,
		PriceCents:      999,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var booksByAuthor1 = map[string]bookstore2.Book{
	"Book1": {
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
var booksByAuthor2 = map[string]bookstore2.Book{
	"Book1": {
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
	"Book2": {
		ID:              "abcde123",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondō"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          30,
		Serie:           1,
		PriceCents:      1199,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var booksByAuthor3 = map[string]bookstore2.Book{
	"Book1": {
		ID:              "abcde345",
		Title:           "Go BootCamp",
		Authors:         []string{"Matt Aimonetti", "X", "Y", "Z"},
		Description:     "Everything you need to know to get started with Go",
		Copies:          40,
		Serie:           1,
		PriceCents:      999,
		DiscountPercent: 0.05,
		Featured:        true,
	},
}

var case1 = "The book Spark Joy with ID abcde123 has a cost of 1199\nThe book Spark Joy with ID abcde123 has a cost of 1199\nThe book Go BootCamp with ID abcde345 has a cost of 999"
var case2 = ""
var case3 = "The book Spark Joy with ID abcde123 has a cost of 1199"

package main

import (
	"bookstore"
	"fmt"
)

func main() {
	// books := []bookstore.Book{}
	// books := []bookstore.Book{
	// 	{
	// 		ID:              "abcde123",
	// 		Title:           "Spark Joy",
	// 		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
	// 		Description:     "A tiny, cheerful Japanese woman explains tidying.",
	// 		Copies:          30,
	// 		Serie:           1,
	// 		PriceCents:      1199,
	// 		DiscountPercent: 0.05,
	// 		Featured:        true,
	// 	},
	// 	{
	// 		ID:              "abcde123",
	// 		Title:           "Spark Joy",
	// 		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
	// 		Description:     "A tiny, cheerful Japanese woman explains tidying.",
	// 		Copies:          30,
	// 		Serie:           1,
	// 		PriceCents:      1199,
	// 		DiscountPercent: 0.05,
	// 		Featured:        true,
	// 	},
	// 	{
	// 		ID:              "abcde123",
	// 		Title:           "Spark Joy",
	// 		Authors:         []string{"Marie Kondō", "Karla Berlanga"},
	// 		Description:     "A tiny, cheerful Japanese woman explains tidying.",
	// 		Copies:          30,
	// 		Serie:           1,
	// 		PriceCents:      1199,
	// 		DiscountPercent: 0.05,
	// 		Featured:        true,
	// 	},
	// }

	// new_books := bookstore.GetAllBooks(books)
	// a := make([]bookstore.Book, 0)
	// b := make([]bookstore.Book, 0)
	// fmt.Println(reflect.DeepEqual(books, new_books))
	// fmt.Println(reflect.DeepEqual(a, b))
	// fmt.Println(bookstore.GetAllBooks(nil))
	// for i := 0; i < 10; i++ {
	fmt.Println(bookstore.NewID())
	// }
}

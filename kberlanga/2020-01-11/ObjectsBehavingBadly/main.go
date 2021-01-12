package main

import (
	"fmt"

	"../ObjectsBehavingBadly/bookstore3"
)

func main() {
	b := &bookstore3.Book{
		PriceCents:      100,
		DiscountPercent: 25,
	}

	net := b.NetPrice()
	fmt.Println(net)
}

package main

import (
	"fmt"
	"mytypes"
	"strings"
)

func main() {
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	fmt.Println(sb.Len())
	fmt.Println(sb.String())

	var mb mytypes.MyBuilder
	fmt.Println(mb.Hello())
	// fmt.Println(mb.Len())
	fmt.Println(mb.Contents.WriteString("Hello, Gophers!"))
	fmt.Println(mb.Contents.Len())

	var su mytypes.StringUppercaser
	// fmt.Println(su.ToUpper())
	// fmt.Println(su.Len())
	fmt.Println(su.WriteString("Congrats, we did it!"))
	fmt.Println(su.ToUpper())
}

package main

import (
	"fmt"
	"strings"

	"code.google.com/p/go-tour/wc"
)

func main() {
	fmt.Println("----------------- Arrays -----------------")
	Arrays()
	fmt.Println("----------------- Slices -----------------")
	Slices()
	fmt.Println("----------------- Range -----------------")
	Range()
	fmt.Println("----------------- Exercise -----------------")
	Exercise()
	fmt.Println("----------------- Maps -----------------")
	Maps()
	fmt.Println("----------------- Exercise 2 -----------------")
	wc.Test(WordCount)
}

/*################################################################################################*/
/*=== arrays ===*/
/*################################################################################################*/
/*
- [n]T is an array of n values of type T
*/

func Arrays() {
	/* array with limit length */
	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"

	/* use an ellipsis to use an implicit length when you pass the values */
	b := [...]string{"hello", "world!", "is", "an", "array", "without", "specific", "length"}
	fmt.Println("++++++++ Ellipsis ++++++++ ")
	fmt.Println(b)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", b)

	/* multidimensional arrays */
	var c [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Println("++++++++ Multidimensional ++++++++ ")
	fmt.Printf("%q\n", c)
}

/*################################################################################################*/
/*=== slices ===*/
/*################################################################################################*/
/*
- slices is like a multidimensional arrays to give a more genera, powerful and convenient interface
- []T is a slice with elements of type T
*/

func Slices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)
	/* a slice of the elements from lo through hi-1, inclusive */
	fmt.Println(p[0:3])
	fmt.Println(p[:3])
	fmt.Println(p[3:])

	/* Making a slice */
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q\n", cities)

	/* Appending to slice */
	names := []string{}
	names = append(names, "Karla")
	names = append(names, "Adam", "Ryan")
	names = append(names, "John")
	other_names := []string{"Harry", "Justin"}
	names = append(names, other_names...)
	fmt.Printf("%q\n", names)

	/* Slice length */
	fmt.Printf("Length of names slice: %v\n", len(names))
}

/*################################################################################################*/
/*=== range ===*/
/*################################################################################################*/
func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	cities := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}

/*################################################################################################*/
/*=== exercise ===*/
/*################################################################################################*/
/*
- Given a list of names, you need to organize each name within a slice based on its length.
*/
func Exercise() {
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha",
		"Addison", "Natalie", "Mia", "Alexis"}

	organize := map[int][]string{}
	for _, name := range names {
		lon := len(name)
		if val, ok := organize[lon]; ok {
			val = append(val, name)
			organize[lon] = val
		} else {
			organize[lon] = []string{name}
			fmt.Println(organize)
		}
	}
	for key, value := range organize {
		fmt.Printf("Names with length %d: %s\n", key, value)
	}
}

/*################################################################################################*/
/*=== maps ===*/
/*################################################################################################*/
/*
- maps is like dictionaries or hashes
*/

type Vertex struct {
	Lat, Long float64
}

func Maps() {
	celebs := map[string]int{
		"Nicolas Cage":       50,
		"Selena Gomez":       21,
		"Jude Law":           41,
		"Scarlett Johansson": 29,
	}

	fmt.Printf("%#v\n", celebs)

	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}
	m["Other"] = Vertex{37.42202, -122.08408}
	fmt.Println(m)
	delete(m, "Other")
	fmt.Println(m)

}

/*################################################################################################*/
/*=== exercise 2 ===*/
/*################################################################################################*/
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}

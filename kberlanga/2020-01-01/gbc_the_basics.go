package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	Hello()
	fmt.Println("-------- Variables --------")
	Variables()
	fmt.Println("-------- Constants --------")
	Constants()
	fmt.Println("-------- Packages --------")
	Packages()
	fmt.Println("-------- Functions --------")
	fmt.Println(add(42, 13))
	region, continent := location("Santa Monica")
	fmt.Printf("Karla lives in %s, %s \n", region, continent)
	// fmt.Println("-------- Pointers --------")
	// Pointers()
	fmt.Println("-------- Mutability --------")
	Mutability()
}

func Hello() {
	fmt.Println("Hello, World!")
}

func Variables() {
	/* list of variables */
	var (
		name, location string
		age            int
	)

	/* can also declared one by one
	var name string
	var location string
	var age int
	*/

	/* can also can iclude initializers one per variable
	var (
		name     string = "Prince Oberyn"
		age      int    =  32
		location string = "Dorne"
	)
	*/

	name, location = "Prince Oberyn", "Dorne"
	age = 32
	fmt.Printf("%s (%d) of %s \n", name, age, location)

	/* a variable can contain any type, including, functions */
	action := func() {
		fmt.Println("A variable is also a function!")
	}

	action()
}

func Constants() {

	/*
		- constants area declared like variables, but with const keyword
		- constants can only be character, string, boolean or numeric values
		- cannot be declared using the := syntax
	*/

	const Pi = 3.14

	/* list of const */
	const (
		Truth = false
		Big   = 1 << 62
		Small = Big >> 61
	)

	fmt.Println("Pi", Pi)
	fmt.Println("Truth", Truth)
	fmt.Println("Big", Big)
}

func Packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*################################################################################################*/
/*=== functions ===*/
/*################################################################################################*/
/*
- a function can take zero or more typed arguments
- the type comes after the variable
- functions can be defined and return any numbers of values that always are typed
*/
/* this function returns a int value */
func add(x int, y int) int {
	return x + y
}

/* and this function returns two strings values can also be named and act just like variables */
func location(city string) (region, continent string) {
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

/*################################################################################################*/
/*=== pointers ===*/
/*################################################################################################*/

/*
- to get the pointer of a value, use the & symbol in front of the value
- to dereference a pointer, use the * symbol
*/

func Pointers() {
	client := &http.Client{}
	resp, err := client.Get("http://gobootcamp.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Resp %v\n", resp)
}

/*################################################################################################*/
/*=== mutability ===*/
/*################################################################################################*/
type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

// func OutMutability() {
// 	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
// 	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
// 	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
// }

func Mutability() {
	artist := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", artist.Name, newRelease(artist))
	fmt.Printf("%s has a total of %d songs\n", artist.Name, artist.Songs)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("----------------- Methods -----------------")
	u := User{"Matt", "Aimonetti"}
	fmt.Println(u.Greeting())
	fmt.Println("----------------- Type aliasing -----------------")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println("----------------- Method receivers -----------------")
	u2 := &User{"Matt", "Aimonetti"}
	fmt.Println(u2.GreetingPointer())
}

/*################################################################################################*/
/*=== methods ===*/
/*################################################################################################*/
/*
- Go isn't an Object Oriented Programming Language
- types and methods allow an object-oriented style of programming
- instaed inheritante, Go has a concept of interface
- Go doesn't have classes but instead can define methods on struct types
- Method is a function on an instance of an object
*/

/* represent an user */
type User struct {
	FirstName, LastName string
}

/* method */
func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

/*################################################################################################*/
/*=== code organization ===*/
/*################################################################################################*/
/*
A good organization, can be:
1. package to import
2. list of constants
3. list of variables
4. main type(s) for the file
5. list of functions
6. list of methods
*/

/*################################################################################################*/
/*=== type aliasing ===*/
/*################################################################################################*/
/*
- To define methods on a type you don’t “own”, you need to define an alias for the type you want to extend
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*################################################################################################*/
/*=== method receivers ===*/
/*################################################################################################*/
/*
- there are two reasons to use a pointer receiver:
1. to avoid copying the value on each method call
2. More efficient if the value type is a large struct
The other reason why you might want to use a pointer is so that the method can modify the value that its receiver points to.
*/

func (u *User) GreetingPointer() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("----------------- Interfaces -----------------")
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))
	fmt.Println("----------------- Interfaces are satisfied implicitly -----------------")
	// var w Writer

	// os.Stdout implements Writer
	// w = os.Stdout

	// fmt.Fprintf(w, "hello, writer\n")

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------- Exercise -----------------")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	planets := map[Planet]float64{
		"Mercury": PeroidToSeconds(0.2408467),
		"Venus":   PeroidToSeconds(0.61519726),
		"Earth":   math.Pow(31557600, -1),
		"Mars":    PeroidToSeconds(1.8808158),
		"Jupiter": PeroidToSeconds(11.862615),
		"Saturn":  PeroidToSeconds(29.447498),
		"Uranus":  PeroidToSeconds(84.016846),
		"Neptune": PeroidToSeconds(164.79132),
	}
	fmt.Println(planets["Venus"])

}

type Planet string

func PeroidToSeconds(period float64) float64 {
	return math.Pow(31557600*period, -1)
}

/*################################################################################################*/
/*=== interfaces ===*/
/*################################################################################################*/
/*
- an interface type is defined by a set of methods
*/

type User struct {
	FirstName, LastName string
}

type Customer struct {
	Id       int
	FullName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

/*################################################################################################*/
/*=== interfaces are satisfied implicitly ===*/
/*################################################################################################*/
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

/*################################################################################################*/
/*=== errors ===*/
/*################################################################################################*/
type error interface {
	Error() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*################################################################################################*/
/*=== exercise ===*/
/*################################################################################################*/
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

package main

import (
	"fmt"
	"time"
)

func main() {
	foo := map[string]interface{}{
		"Matt":  42,
		"Karla": 23,
	}
	fmt.Println("-------- Type assertion --------")
	timeMap(foo)
	fmt.Println(foo)
	fmt.Println("-------- Structs --------")
	event := Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	}
	fmt.Println(event)
	fmt.Println(p, q, r, s)
	fmt.Println("Acessing fields of struct using the dor notation")
	fmt.Printf("Event on %s, location (%f, %f)\n", event.Date, event.Lat, event.Lon)
	fmt.Println("-------- Initializing --------")
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(*x == *y)
	fmt.Println("-------- Exercise --------")
	player := &Player{
		User:   &User{Id: 1, Name: "Karla", Location: "MX"},
		GameId: 123,
	}
	fmt.Println(*player.User, player.GameId)
	fmt.Println(player.Greetings())
}

/*################################################################################################*/
/*=== type assertion ===*/
/*################################################################################################*/
func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

/*################################################################################################*/
/*=== structs ===*/
/*################################################################################################*/
type Bootcamp struct {
	// Latitude of the event
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}  // has type Point
	q = &Point{1, 2} // has type *Point
	r = Point{X: 1}  // Y:0 is implicit
	s = Point{}      // X:0 and Y:0
)

/*################################################################################################*/
/*=== excersise ===*/
/*################################################################################################*/
type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

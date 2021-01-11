package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------------------- Goroutines -----------------------")
	go say("world")
	say("hello")
	fmt.Println("----------------------- Channels -----------------------")
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fmt.Println("+++++++ Buferred channels +++++++")
	d := make(chan int, 2)
	d <- 1
	d <- 2
	// d <- 3
	d3 := func() { d <- 3 }
	go d3()
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println("----------------------- Range and close -----------------------")
	e := make(chan int, 10)
	go fibonacci(cap(e), e)
	for i := range e {
		fmt.Println(i)
	}
	fmt.Println("----------------------- Select -----------------------")
	f := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-f)
		}
		quit <- 0
	}()
	fibonacci2(f, quit)

	fmt.Println("+++++++ default +++++++")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*################################################################################################*/
/*=== goroutines ===*/
/*################################################################################################*/
/*
- a goroutine is a lightweight thread managed by Go runtime
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*################################################################################################*/
/*=== channels ===*/
/*################################################################################################*/
/*
- a channel is a typed condiut throught which you can send and recreive values with the channel operator <-
- the data flows in the direction of the arrow
- the channles must be created before use `ch := make(chan int)`
*/

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

/*################################################################################################*/
/*=== range and close ===*/
/*################################################################################################*/
/*
- a sender can close a channel to indicate that no more values will be sent.
- only the sender should close a channel, never de receiver.
- channels aren't like files; you don't usually need to close them
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*################################################################################################*/
/*=== select ===*/
/*################################################################################################*/
/*
- the select stachment lets a goroutine wait on multiple communications operations
*/

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("----------------- IfStatement -----------------")
	IfStatement()
	fmt.Println("----------------- ForLoop -----------------")
	ForLoop()
	fmt.Println("----------------- SwitchCase -----------------")
	SwitchCase()
	fmt.Println("----------------- Exercise -----------------")
	Exercise()
}

/*################################################################################################*/
/*=== if ===*/
/*################################################################################################*/
func IfStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

/*################################################################################################*/
/*=== for ===*/
/*################################################################################################*/
func ForLoop() {
	/* classic for */
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/* for without pre/post statements */
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	/* for as while */
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)
}

/*################################################################################################*/
/*=== switch case ===*/
/*################################################################################################*/
/*
interesting things
- you can only compare values of the same type
- you can set an optional default statement to be executed if all the other fail
- you can use an expression in the case statement, for instance you can calculate a value to use in this case
- you can have multiple values in case statement
- you can execute all the following statements after a match using the `fallthrough` statement
- you can use a break statement inside your matched statement to exit the switch processing
*/
func SwitchCase() {
	now := time.Now().Unix()
	fmt.Println("Now  time is ", now)
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	/* you can use an expression in the case statement, for instance you can calculate a value to use in this case */
	num := 3
	v := num % 2
	switch v {
	case 0:
		fmt.Println("even")
	case 3 - 2:
		fmt.Println("odd")
	}

	/* you can have multiple values in case statement */
	score := 6
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}

	/* you can execute all the following statements after a match using the `fallthrough` statement */
	n := 7
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}
}

/*################################################################################################*/
/*=== exercise ===*/
/*################################################################################################*/
/*
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can’t get more than 10 coins.
Print a map with each user’s name and the amount of coins distributed. After distributing all the coins, you should have 2 coins left.
*/

func Exercise() {
	total_coins := 50
	users := map[string]int{
		"Matthew":   0,
		"Sarah":     0,
		"Augustus":  0,
		"Heidi":     0,
		"Emilie":    0,
		"Peter":     0,
		"Giana":     0,
		"Adriano":   0,
		"Aaron":     0,
		"Elizabeth": 0,
	}

	for name := range users {
		coins_name := CountVowels(name)
		fmt.Println(total_coins)
		fmt.Println(name)
		fmt.Println(coins_name)
		if total_coins-coins_name < 3 {
			users[name] = coins_name
			total_coins -= coins_name
			break
		}
		if coins_name < 11 {
			users[name] = coins_name
			total_coins -= coins_name
		} else {
			users[name] = 10
			total_coins -= 10
		}
	}

	fmt.Println(users)
	fmt.Println("Coins left: ", total_coins)
}

func CountVowels(str string) (total int) {
	total = 0
	for _, letter := range str {
		switch string(letter) {
		case "a", "A":
			total += 1
		case "e", "E":
			total += 1
		case "i", "I":
			total += 2
		case "o", "O":
			total += 3
		case "u", "U":
			total += 4
		}
	}
	return total
}

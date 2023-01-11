package main

import "fmt"

func main() {
	// Operator aritmatik
	var discounts float64 = 0.1 * 10000
	fmt.Println("Total discont:", discounts)

	var (
		leftSide  = 13
		rightSide = 15
	)
	balanced := (leftSide == rightSide)
	fmt.Println("Balanced:", balanced)
}

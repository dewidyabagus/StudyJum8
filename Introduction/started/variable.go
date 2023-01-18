package main

import "fmt"

func main() {
	var (
		roomNumber  uint8 // Camel Case
		floorNumber uint8 = 13
	)

	roomNumber = 135
	fmt.Println("Room Number:", roomNumber)

	// var floorNumber uint8 = 13
	fmt.Println("Floor Number:", floorNumber)

	firstName, middleName, lastName, age, merried := "Johs", "Aliando", "Wick", uint64(12), true
	fmt.Println("Full Name:", firstName, middleName, lastName, age, merried)

	// const pi float64 = 3.14
	// const pi2 = 3.14
	const (
		pi  float64 = 3.14
		pi2         = 3.14
	)
}

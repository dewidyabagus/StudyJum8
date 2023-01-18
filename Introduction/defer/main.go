package main

import "fmt"

func main() {
	// defer fmt.Println("Defer #1")
	// defer fmt.Println("Defer #2")
	// defer fmt.Println("Defer #3")

	// Eksekusi defer LIFO (Last In First Out)
	// OUTPUT: Defer #3
	// OUTPUT: Defer #2
	// OUTPUT: Defer #1

	// client := OpenSession()
	// defer client.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover panic by defer:", r)
		}
	}()

	panic("Testing")

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recover panic by defer:", r)
	// 	}
	// }()
}

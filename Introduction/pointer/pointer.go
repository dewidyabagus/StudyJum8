package main

import "fmt"

func main() {
	// & Referencing
	// * Dereferencing
	var x int = 100
	fmt.Println("Nilai X adalah", x)

	// var y *int = &x // Long
	y := &x // Short langsung inisialisasi
	fmt.Println("Nilai Y adalah", y)
	fmt.Println("Nilai Y Original adalah", *y)

	*y = 121
	// y = 132 // Invalid: cannot use 132 (untyped int constant) as *int value in assignment

	fmt.Println("Nilai X #2 adalah", x)
	fmt.Println("Nilai Y #2 adalah", *y)
}

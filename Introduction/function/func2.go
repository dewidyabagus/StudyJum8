package main

import (
	"fmt"
	"math"
)

// func calculation(d float64) float64
// func calculation(d float64) (float64, float64) {
// 	area := math.Pi * math.Pow(d/2, 2)

// 	circumrference := math.Pi * d

// 	return area, circumrference
// }

func calculation(d float64) (area float64, circumrference float64) {
	area = math.Pi * math.Pow(d/2, 2)

	circumrference = math.Pi * d

	return
}

func main() {
	var diameter float64 = 14

	area, circumference := calculation(diameter)
	// var area, circumference = calculation(diameter)
	// var (
	// 	area, circumference = calculation(diameter)
	// )
	fmt.Println("Area:", area)
	fmt.Println("Keliling:", circumference)
}

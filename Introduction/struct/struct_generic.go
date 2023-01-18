package main

import "fmt"

type Types interface {
	int | uint | float64
}

// type Divide[T int | uint | float64] struct {
// 	Y T
// 	X T
// }

type Divide[T Types] struct {
	X T
	Y T
}

func (d *Divide[T]) Calculate() T {
	return (d.X / d.Y)
}

func main() {
	divide := Divide[int]{
		X: 6,
		Y: 2,
	}
	fmt.Println("Hasil", divide.Calculate())
}

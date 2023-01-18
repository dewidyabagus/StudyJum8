package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		x int = 6
		y int = 0
	)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Menangkap panic:", r)
		}
	}()

	res, err := divide(x, y)
	if err != nil {
		// log.Println("Error:", err.Error())
		panic(fmt.Sprintf("Error: %s", err.Error()))
	} else {
		fmt.Printf("Pembagian %d / %d = %d \n", x, y, res)
	}

	// Tidak akan kebawah
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("y bernilai 0, pembagian dibatalkan")
	}
	res := x / y
	return res, nil
}

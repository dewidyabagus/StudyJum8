package main

import (
	"errors"
	"fmt"
	"log"
)

func sliceCalc(numbers []int, calc func(slice []int) int) (int, error) {
	val := calc(numbers)

	// val digunakan sebagai acuan blok kode dibawah
	// val.....
	if len(numbers) == 0 {
		return 0, errors.New("slice tidak boleh kosong")
	}

	return val, nil
}

func main() {
	var numbers = []int{3, 2, 5, 6, 4, 4, 5}
	fmt.Println("Input list numbers :", numbers)

	var min, err = sliceCalc(numbers, func(slice []int) int {
		var min = slice[0]

		for _, val := range slice {
			if val < min {
				min = val
			}
		}

		return min
	})
	if err != nil {
		log.Fatalln("Error min:", err.Error())
	}
	fmt.Println("Min value:", min)

	var max, err1 = sliceCalc(numbers, func(slice []int) int {
		var max = slice[0]

		for _, val := range slice {
			if val > max {
				max = val
			}
		}

		return max
	})
	if err1 != nil {
		log.Fatalln("Error Max:", err.Error())
	}
	fmt.Println("Max value:", max)
}

package main

import (
	"errors"
	"fmt"
	"log"
)

func sliceCalc[T int | float32 | uint](numbers []T, calc func(slice []T) T) (T, error) {
	if len(numbers) == 0 {
		return 0, errors.New("slice tidak boleh kosong")
	}

	return calc(numbers), nil
}

// func sliceCalcFloat32(numbers []float32, calc func(slice []float32) float32) (float32, error) {
// 	if len(numbers) == 0 {
// 		return 0, errors.New("slice tidak boleh kosong")
// 	}

// 	return calc(numbers), nil
// }

// func sliceCalcUint(numbers []uint, calc func(slice []uint) uint) (uint, error) {
// 	if len(numbers) == 0 {
// 		return 0, errors.New("slice tidak boleh kosong")
// 	}

// 	return calc(numbers), nil
// }

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
}

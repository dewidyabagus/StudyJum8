package main

import (
	"fmt"
)

func averageCalculation(numbers ...float32) float32 {
	numbers[0] = 2
	fmt.Printf("Alamat memory numbers: %p", numbers)
	var total float32

	for _, num := range numbers {
		total += num
	}

	var avg = total / float32(len(numbers))

	return avg
}

func modifySlice(numbers []float32) {
	// numbers[0] = 2
	numbers = append(numbers, 2)
}

func main() {
	// var number = []float32{1, 3, 4, 5, 6}
	// var avg = averageCalculation(number...)
	// fmt.Println(avg)

	var nilai float32 = 1
	fmt.Printf("Alamat memory nilai: %p \n", &nilai)

	var avg = averageCalculation(nilai)
	fmt.Println(avg)
	fmt.Println("Nilai:", nilai)

	numbers := []float32{1}
	modifySlice(numbers)
	fmt.Println("Numbers:", numbers)
	// var number string = "1"
	// var numberInt, _ = strconv.Atoi(number)

	// var number uint64 = 257
	// var numb uint8 = uint8(number)

	// fmt.Println("Number:", numb)
}

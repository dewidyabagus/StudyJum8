package main

import "fmt"

// func getMinMax ....
func main() {
	// var getMinMax = func(numbers []int) (min, max int) {
	// 	if len(numbers) == 0 {
	// 		return
	// 	}

	// 	min, max = numbers[0], numbers[0]

	// 	// for (index), (val) := range (slice/ array/ map)
	// 	for _, val := range numbers {
	// 		if val < min {
	// 			min = val
	// 		} else if val > max {
	// 			max = val
	// 		}
	// 	}

	// 	return
	// }

	// // Pemanggilan function
	var numbers = []int{2, 3, 4, 3, 2, 3}
	// var min, max = getMinMax(numbers)
	// fmt.Println("Min:", min)
	// fmt.Println("Max:", max)

	var min, max = func(numbers []int) (min, max int) {
		if len(numbers) == 0 {
			//
			return
		}

		min, max = numbers[0], numbers[0]

		// for (index), (val) := range (slice/ array/ map)
		for _, val := range numbers {
			if val < min {
				min = val
			} else if val > max {
				max = val
			}
		}

		return
	}(numbers)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

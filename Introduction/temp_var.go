package main

// import "fmt"
// import "math"

import (
	"fmt"
	"math"
)

func main() {
	var edge float64 = 30

	var message = "from main function"
	if cubeVolume := math.Pow(edge, 3); cubeVolume > 10_000 {
		fmt.Println("Lerge volume cube. volume:", cubeVolume)
		var message = "from block if cubeVolume > 10_000"
		fmt.Println(message)
		// if i := 1; i > 5 {
		// 	fmt.Println(i)
		// 	if n := 3; n > 6 {
		// 		fmt.Println(n)
		// 	}
		// 	fmt.Println(n)
		// }
		// fmt.Println(i)
	} else if cubeVolume >= 5000 {
		fmt.Println("Mediun volume cube. volume:", cubeVolume)

	} else {
		fmt.Println("Small volume cube. volume:", cubeVolume)
	}

	fmt.Println(message)
}

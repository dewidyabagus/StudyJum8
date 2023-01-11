package main

import "fmt"

func main() {
	// var i = 0 // Inisialisasi
	// for (kondisi) {
	//    (block code)
	//
	//    (iterasi)
	//}
	var i int = 0

	for i < 5 {
		fmt.Println("Cetak perulangan i ke:", i+1)

		i++
	}

	var n = 0
	for ; n < 5; n++ {
		fmt.Println("Cetak perulangan n ke:", n+1)
	}
}

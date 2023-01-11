package main

import "fmt"

func main() {
	lists := [5]int{1, 2, 3, 4, 5}

	// Penggunaan key for - range
	// for (index), (element value) := range (array/ slice/ map) {
	//    kode yang akan dieksekusi
	// }
	for i, val := range lists {
		fmt.Println("Cetak nilai index ke:", i, "nilai:", val)
	}
}

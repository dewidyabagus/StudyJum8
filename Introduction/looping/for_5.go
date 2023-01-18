package main

import "fmt"

func main() {

	var i = 1
	for {
		if i == 5 {
			break
		}
		fmt.Println("Cetak perulangan ke :", i)

		i++
	}
}

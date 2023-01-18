package main

import "fmt"

func main() {
	var lists = map[string]int{"key1": 5}

	// Kita ingin mencari key dengan nama key1 apakah ada dalam lists
	// val, ok := lists["key1"]
	// if !ok {
	// 	fmt.Println("Key1 tidak ada")
	// } else {
	// 	fmt.Println("Key1 ada dengan nilai", val)
	// }

	if val, ok := lists["key1"]; !ok {
		fmt.Println("Key1 tidak ada")
	} else {
		fmt.Println("Key1 ada dengan nilai", val)
	}
}

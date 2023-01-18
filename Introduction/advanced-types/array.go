package main

import "fmt"

func main() {
	// Long Desclaration
	// var hobbies [3]string
	// hobbies[0] = "bermain"
	// hobbies[1] = "bersepeda"
	// hobbies[2] = "bermain bola"
	// fmt.Println("Hobbies:", hobbies)

	// Short Declaration, Inisialisasi nilai
	fullName := [3]string{"Widya", "Ade", "Bagus"}
	fmt.Println("Full Name:", fullName)
	// fullName[0] = "Angga"
	// fmt.Println("Full Name #2:", fullName)
	fmt.Println("Length Full Name:", len(fullName))
	fmt.Println("Capacity Full Name:", cap(fullName))

	var hobbies = [...]string{"Bermain bola", "Bersepeda", "Coding"}
	fmt.Println("Hobbies:", hobbies)

	// var tables = [3][4]string{
	// 	{"A", "B", "C", "D"},
	// 	{"E", "F", "G", "H"},
	// }
	var tables [3][4]string
	tables[0][0] = "A"
	fmt.Println("Array multi dimention:", tables)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Definisikan array
	// var hobbies []string
	// hobbies = []string{}
	// Append
	// hobbies = append(hobbies, "Main Bola")
	// fmt.Println("Hobbies:", hobbies)

	// var hobbies = []string{"Belajar", "membace", "berlari"}
	// fmt.Println("Hobbies:", hobbies)

	// Array
	// var hobbies = [3]string{"Belajar", "Coding", "Bersepeda"}
	// var hobbies1 = hobbies[0:3]
	// var hobbies2 = hobbies[:1]
	// var hobbies3 = hobbies[1:]
	// hobbies3[0] = "Bersepeda"

	// fmt.Println("Hobbies 1:", hobbies1)
	// fmt.Println("Hobbies 2:", hobbies2)
	// fmt.Println("Hobbies 3:", hobbies3)

	// var hobbies = []string{"menulis", "bermain game", "berenang", "coding"}
	// fmt.Println("Length hobbies:", len(hobbies))
	// fmt.Println("Capacity hobbies:", cap(hobbies))
	// fmt.Printf("Alamat memory: %p \n", hobbies)
	// var hobbies1 = hobbies[:]
	// hobbies = append(hobbies, "berselancar")
	// hobbies[0] = "Coding"
	// fmt.Printf("Alamat memory: %p \n", hobbies)

	// fmt.Println("Hobbies 1:", hobbies1)
	// fmt.Println("Hobbies  :", hobbies)

	// Langsung Mengisi element
	// var hobbies = make([]string, 3)
	// hobbies = append(hobbies, "")
	// hobbies = append(hobbies, "")
	// fmt.Println("Length hobbies:", len(hobbies))
	// fmt.Println("Capacity hobbies:", cap(hobbies))

	// var hobbies = make([]string, 0, 7)
	// hobbies = append(hobbies, "Hobby Pertama")
	// fmt.Printf("Alamat memory: %p \n", hobbies)
	// hobbies = append(hobbies, "Hobby Kedua")
	// fmt.Printf("Alamat memory: %p \n", hobbies)
	// fmt.Println("Length hobbies:", len(hobbies))
	// fmt.Println("Capacity hobbies:", cap(hobbies))

	// var names = []string{"bejo", "heru", "achmad"}
	// var names = make([]string, 0, 3)
	// names = append(names, "bejo")
	// fmt.Printf("Alamat memory: %p \n", names)

	// var names2 = names[:]

	// names = append(names, "rudi")
	// names = append(names, "heru") // Maksimal
	// fmt.Printf("Alamat memory: %p \n", names)

	// names = append(names, "Agus")
	// fmt.Printf("Alamat memory: %p \n", names)
	// names2[0] = "achmad"
	// // fmt.Println("Length Names:", len(names))
	// fmt.Println("Names 1:", names)
	// fmt.Println("Names 2:", names2)
	// fmt.Printf("Alamat memory 2: %p \n", names2)
	// fmt.Println("Capacity hobbies:", cap(names))

	var names = make([]string, 3)
	for i := 0; i < len(names); i++ {
		names[i] = strconv.Itoa(i)
	}
	names = make([]string, 3)
	fmt.Println("names:", names)
}

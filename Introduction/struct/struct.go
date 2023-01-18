package main

import (
	"fmt"
	"strings"
)

type Biodata struct {
	FistName    string
	LastName    string
	Age         int
	Hobbies     []string
	DateOfBirth string
}
type NewString string
type NewBiodata Biodata      // Membuat tipe data baru
type NewBiodataRef = Biodata // References struct

// func (b Biodata) FuncName()
func (this Biodata) PrintFullName() {
	fmt.Println("Full name:", this.FistName, this.LastName)
}

// Dereferencing
// func (this *Biodata) AddHobby(hobby string) {
// 	this.Hobbies = append(this.Hobbies, hobby)
// }

// Biasa
func (this Biodata) AddHobby(hobby string) {
	this.Hobbies = append(this.Hobbies, hobby)
	fmt.Printf("Alamat memori hobbies [method]: %p \n", this.Hobbies)
	fmt.Println("Hobbies [method]:", strings.Join(this.Hobbies, ", "))
}

func main() {
	// var biodata Biodata // Long
	// biodata := Biodata{} // Short
	biodata := Biodata{
		FistName:    "Jonh",
		LastName:    "Wick",
		Age:         28,
		Hobbies:     []string{"berenang", "coding"},
		DateOfBirth: "Jakarta",
	}
	biodata.Age = 30

	fmt.Printf("Alamat memori hobbies [luar]: %p \n", biodata.Hobbies)
	// fmt.Println("Biodata:", biodata)
	// OUTPUT: Biodata: {Jonh Wick 30 [berenang coding] Jakarta}

	// fmt.Printf("Biodata: %+v \n", biodata)
	// OUTPUT: Biodata: {FistName:Jonh LastName:Wick Age:30 Hobbies:[berenang coding] DateOfBirth:Jakarta}

	// biodata.PrintFullName()
	// OUTPUT: Full name: Jonh Wick

	// list := []string{"1", "2"}
	// strings.Join(list, ", ") // "1", "2" (string)
	biodata.AddHobby("memancing")
	fmt.Println("Hobbies:", strings.Join(biodata.Hobbies, ", "))
}

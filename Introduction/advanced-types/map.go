package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// var biodata map[string]string
	// biodata = map[string]string{}
	// biodata["name"] = "john wick"
	// fmt.Println("biodata:", biodata)

	// Long Declaration
	// var biodata map[string]string // Inisalisasi zero value nil

	// Short Declaration
	// biodata := map[string]string{}

	// Menggunakan function make()
	// biodata := make(map[string]string{})

	biodata := map[string]string{
		"hobby":  "memancing",
		"alamat": "Kediri Jawa Timur",
		"nama":   "John Wick",
	}
	// biodata1 := map[string]string{"nama": "John Wick", "alamat": "Kediri Jawa Timur", "hobby": "memancing"}
	// fmt.Println("biodata:", biodata)

	// fmt.Println("Name:", biodata["nama"])
	// fmt.Println("Alamat:", biodata["alamat"])
	// fmt.Println("Sekolah:", biodata["sekolah"])

	// val, found := biodata["sekolah"]
	// if !found {
	// 	fmt.Println("Cetak tidak diketemukan untuk key sekolah")
	// }
	// _ = val

	for key, val := range biodata {
		fmt.Println(key, "=>", val)
	}

	body, _ := json.Marshal(biodata)
	fmt.Println(string(body))
}

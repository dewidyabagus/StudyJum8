package main

import "fmt"

func main() {
	// PrintDataType(uint(121))

	var data interface{} = 1

	if data.(bool) {
		fmt.Println("Kondisi benar")
	}
}

// Data type any
func PrintDataType(data interface{}) {
	switch c := data.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("Tipe data signed number :", c)

	case uint, uint8, uint16, uint32, uint64:
		fmt.Println("Tipe data unsigned number:", c)

	case string:
		fmt.Println("Tipe data string:", c)
	}
}

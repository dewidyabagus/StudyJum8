package main

func main() {
	// var grade = "Grade D"

	// switch grade {
	// case "Grade A":
	// 	fmt.Println("Point range 95 - 100")

	// case "Grade B":
	// 	fmt.Println("Point range 80 -94")

	// case "Grade C":
	// 	fmt.Println("Point range 70 - 79")

	// default:
	// 	fmt.Println("Point range 0 - 69")
	// }

	// var points uint8 = 85

	// switch {
	// case points >= 95:
	// 	fmt.Println("Grade A")

	// case points < 95 && points >= 80:
	// 	fmt.Println("Grade B")

	// case points < 80 && points >= 70:
	// 	fmt.Println("Grade C")

	// default:
	// 	fmt.Println("Grade D")
	// }

	var transactionType uint8 = 1

	// 1 - Stock Masuk
	// 2 - Stock Keluar
	// 3 - Penjualan
	// 4 - Stock Opname
	// 5 - Stock Terbuang
	// 6 - Mutasi Stock Masuk
	if transactionType == 1 || transactionType == 4 || transactionType == 6 {
		// Lakukan sesuatu
	}

	switch transactionType {
	case 1, 4, 6:
		// Lakukan sesuatu
	}

	var condition = true
	// if condition // statement Invalid Code
	if condition {
	} // action code
}

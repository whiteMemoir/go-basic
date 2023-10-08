package main

import "fmt"

func main() {
	// Contoh single switch case statement
	time := "am"
	switch time {
	case "am":
		fmt.Println("Pagi")
	case "pm":
		fmt.Println("Malam")
	default:
		fmt.Println("Waktunya salah!")
	}

	// Contoh multi switch case statement
	hari := "senin"

	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("Weekday")
	case "sabtu", "minggu":
		fmt.Println("Weekend")
	default:
		fmt.Println("Hari tidak valid!")
	}
}
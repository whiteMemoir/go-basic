package main

import "fmt"

func main() {
	defer fmt.Println("Program selesai")

	var num1, num2 int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Printf("Hasil penjumlahan: %d\n", sum)
}

package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}
	}()

	result := divide(num1, num2)
	fmt.Println("Hasil pembagian:", result)
}

func divide(a, b int) int {
	if b == 0 {
		panic("Pembagian oleh 0 tidak diizinkan")
	}
	return a / b
}

package main

import "fmt"

func main() {
	/*STRING*/
	//Contoh tipe data String
	var txt string = "kutip dua"

	fmt.Println(txt)
	fmt.Printf("Value: %v, tipe data %T\n", txt, txt)

	var txt2 string = `
	Ini adalah "teks".
	Dan ini dari 'baris baru'.
	Oke...
	`

	fmt.Println(txt2)

	/*INTEGER*/
	// Contoh tipe data integer
	var a int16 = 3000
	var age uint8 = 33
	var bigNumber int64 = -987654321987654321

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", age)
	fmt.Printf("%d\n", bigNumber)

	/*FLOAT*/
	// Contoh tipe data float
	var x float64 = 3.14

	fmt.Printf("%f\n", x)
	fmt.Printf("%.1f\n", x)
	fmt.Printf("%.2f\n", x)
	fmt.Println(x)

	/*BOOLEAN*/
	// Contoh tipe data boolean
	var isCompleted bool
	var isMarried bool = true

	fmt.Printf("%t\n", isCompleted)
	fmt.Printf("%t\n", isMarried)
}
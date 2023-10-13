package main

import "fmt"

func main(){
	// Latihan 1
	/** 
			Silahkan buat bentuk seperti ini :
			*****
			****
			***
			**
			*
	*/
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// Latihan 2
	/** 
			Silahkan buat bentuk seperti ini :
			*
			**
			***
			****
			*****
	*/
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// Latihan 3
	/**
			Silahkan buat bentuk Noo Bee (panjang : 50) dengan kondisi berikut :
			Jika habis dibagi 3 : Noo
			Jika habis dibagi 5 : Bee
			Jika habis dibagi 3 dan 5 : Noo Bee

			Contoh :
			1, 2, Noo, 4, Bee, Noo, 7, 8, Noo, Bee, 11, Noo, 13, 14, NooBee, 16, 17, Noo, 19, Bee, ... n
	*/
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("NooBee, ")
		} else if i%3 == 0 {
			fmt.Print("Noo, ")
		} else if i%5 == 0 {
			fmt.Print("Bee, ")
		} else {
			fmt.Printf("%d, ", i)
		}
	}
}
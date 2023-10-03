package main

import "fmt"

var x int = 100

// y := 2000 // Tidak bisa diluar fungsi

func main() {
	/*membuat variable*/
	var myName string = "Saya"
	fmt.Println(myName)

	var yourName string
	yourName = "Kamu"
	fmt.Println(yourName)

	var firstName = "Nama"
	fmt.Println(firstName)

	iniNama := "Nama-nama"
	fmt.Println(iniNama)

	myAge := 19
	fmt.Println(myAge)

	// print var global
	fmt.Println(x)
	// fmt.Println(y)


	/*Deklarasi variabel tanpa nilai awal*/
	var a string
	var b int
	var c bool

	fmt.Println(a) //menghasilkan output nilai default yaitu ""
	fmt.Println(b) //menghasilkan output nilai default yaitu 0
	fmt.Println(c) //menghasilkan output nilai default yaitu false


	/*Deklarasi Multiple Variabel*/
	var d, e, f, g int = 1, 2, 3, 4

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var h, i, j = 0.5, "coba", true

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	/*Deklarasi Multiple Variabel dalam Blok*/
	var (
		earlyName string = "String"
		lastname				= "Coba"
		angka int = 325
	)

	fmt.Println(earlyName)
	fmt.Println(lastname)
	fmt.Println(angka)

}
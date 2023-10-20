package main

import "fmt"

func modifiedString(strings map[string]string) string {
	return "Mobil " + strings["name"] + " berwarna " + strings["color"]
}

func printString(str string) {
	fmt.Println(str)
}

func main(){
	var car map[string]string
	car = make(map[string]string)
	car["name"] = "BWM"
	car["color"] = "Black"

	output := modifiedString(car)
	printString(output)

	// fmt.Println(output)

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// output => Mobil BMW berwarna Black

}
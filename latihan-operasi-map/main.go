package main

import (
	"fmt"
)

func main() {
	fruits := map[string]int{
    "Apple": 10,
    "Banana": 15,
    "Orange": 8,
    "Grape": 20,
	}

	fmt.Println("Sebelum ditambah/hapus: ", fruits)

	fruits["Mango"] = 12 
	fruits["Strawberry"] = 7

	fmt.Println("Sebelum menambahkan Mango dan Strawberry: ", fruits)

	delete(fruits, "Orange")

	fmt.Println("Setelah menghapus buah Orange: ", fruits)

	findKey := "Apple"
	_, isExist := fruits[findKey]
	if isExist {
		fmt.Println("Jumlah apel yang tersedia adalah ", fruits["Apple"])
	} else {
		fmt.Println("Buah apel tidak ditemukan" )
	}

	fruitsSlice := []map[string]interface{}{
		{"nama_buah":"Banana", "qty": 15,},
		{"nama_buah":"Orange", "qty": 20,},
		{"nama_buah":"Grape", "qty": 12,},
		{"nama_buah":"Strawberry", "qty": 7,},
		{"nama_buah":"Apple", "qty": 10,},
	}

	for _, slice := range fruitsSlice {
		fmt.Println(slice["nama_buah"].(string) + ": " + fmt.Sprint(slice["qty"]))
	}
}
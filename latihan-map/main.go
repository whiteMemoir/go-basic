package main

import "fmt"

func main() {
	contactList := make(map[string]int)

	contactList["Alice"] = 1234567890
	contactList["Bob"] = 9876543210

	fmt.Println("Semua Kontak", contactList)
	fmt.Println("Nomor telepon Alice:", contactList["Alice"])

	contactList["Charlie"] = 5555555555

	fmt.Println("Setelah menambah kontak Charlie:", contactList)

	delete(contactList, "Bob")

	fmt.Println("Setelah hapus kontak Bob:", contactList)
	fmt.Println("Data Kontak: ")
	
	for key, value := range contactList {
		fmt.Println(key, ":", value)
	}
}
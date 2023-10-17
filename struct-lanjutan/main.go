package main

import "fmt"

type Parent struct {
	Nama	string
	Umur 	int
}

type Student struct {
	OrangTua	Parent  //embedded struct
	Nama			string
	Umur			int
	Kelas			string
}

//Deklarasi slice of struct
type Employee struct {
	Name				string
	Department	string
}

type Employees []Employee

func main() {
	// Penerapan embedded Struct
	parent1 := Parent{
		Nama: "Budi",
		Umur: 45,
	}

	student1 := Student{
		OrangTua: parent1,
		Nama: "Andi",
		Umur: 14,
		Kelas: "7A",
	}

	fmt.Println(student1)
	fmt.Printf("Orang tua %s bernama %s\n", student1.Nama, student1.OrangTua.Nama)

	student2 := Student{
		OrangTua: Parent{
			Nama: "Udin",
			Umur: 45,
		},
		Nama: "Ida",
		Umur: 12,
		Kelas: "6B",
	}

	fmt.Println(student2)
	fmt.Printf("Orang tua %s bernama %s\n", student2.Nama, student2.OrangTua.Nama)

	// Penerapan slice of struct
	var employees Employees
	fmt.Println(employees)

	var emp1 Employee
	emp1.Name = "Ahmad"
	emp1.Department = "Sales"

	// emp2 := Employee{
	// 	Name: "Habib",
	// 	Department: "Admin",
	// }


}
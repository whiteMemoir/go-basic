package main

import "fmt"

// Parameter Pointer
func changeName(name *string, newName string)  {
	fmt.Println("Change name from", *name, "to", newName)
	*name = newName
}

// Pointer pada Method
type Car struct {
	Name		string
	Color		string
}

func (c *Car) setName(newName string) {
	fmt.Println("Change name from", c.Name, "name to =>", newName)
	c.Name = newName
}

func (c Car) changeName2(newName string) {
	fmt.Println("Change name from", c.Name, "name to =>", newName)
	c.Name = newName
	fmt.Println("di dalam fungsi changeName2", c.Name)
}

func main() {
	nama := "HeeHee"
	namaPointer := &nama

	fmt.Println(nama)
	fmt.Println(namaPointer) // Mengambil alamat memori
	fmt.Println(*namaPointer) // Mengambil value

	var ptr *int
	fmt.Println(ptr)

	//Contoh penerapan pointer
	var num1 int = 5
	var num2 *int = &num1

	fmt.Println("===== Nilai Awal =====")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)


	fmt.Println("===== Nilai Setelah num1 diubah =====")
	num1 = 6
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num2 diubah =====")
	*num2 = 10
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	// Parameter Pointer
	myName := "haaHaa"
	fmt.Println("Nama awal:", myName)

	changeName(&myName, "haaHaaID")
	fmt.Println("Nama setelah diubah:", myName)

	// Pointer pada Method
	car := Car {
		Name: "Toyota",
		Color: "Black",
	}

	car.setName("Honda")
	fmt.Println(car)

	car.changeName2("Mitsubishi")
	fmt.Println(car) // Tidak berubah karena karena ga pake pointer

	
}
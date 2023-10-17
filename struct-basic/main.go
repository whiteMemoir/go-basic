package main

import "fmt"

// Membuat Struct
type Fruit struct {
	Name 		string
	Weight	int
	Price		int
}

func main() {
	//Penerapan struct
	var fruit Fruit
	fruit.Name = "Apple"
	fruit.Weight = 1

	fmt.Println(fruit)
	fmt.Println(fruit.Name)
	fmt.Println(fruit.Weight)

	//Penerapan Struct 2
	fruit1 := Fruit{
		Name: "Mango",
		Weight: 2,
	}

	fmt.Println(fruit1)
	fmt.Println(fruit1.Name)
	fmt.Println(fruit1.Weight)
	
}


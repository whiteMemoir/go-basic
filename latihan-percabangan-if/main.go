package main

import "fmt"

func main() {
	num1 := 4
	num2 := 5

	isOddOrEven(num1)
	isOddOrEven(num2)
}

func isOddOrEven(num int)  {
	if num % 2 == 0 {
		fmt.Printf("%d Merupakan bilangan genap\n", num)
	} else if num % 2 == 1 {
		fmt.Printf("%d Merupakan bilangan ganjil\n", num)
	} else {
		fmt.Printf("Input %d tidak valid!", num)
	}
}
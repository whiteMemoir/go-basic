package main

import "fmt"

func main() {
	// Operator Aritmatika
	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 5)
	fmt.Println(4 / 2)
	fmt.Println(5 % 2)
	fmt.Println("================================================================")

	var num1 int = 10
	var num2 int = 2
	str1 := "Ini"
	str2 := "Budi"

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	fmt.Println(str1 + " " + str2)

	// Operator Increment dan Decrement
	i := 0
	// Ga bisa ++i
	i++
	fmt.Println(i)

	j := 5
	j--
	fmt.Println(j)

	// Operator Penugasan
	var a, b, c, d, e = 5, 10, 15, 20, 25

	var x = 5
	a += 5
	b -= 5
	c *= 5
	d /= 5
	e %= 5

	fmt.Printf("Nilai x: %d\n", x)
	fmt.Printf("Nilai a: %d\n", a)
	fmt.Printf("Nilai b: %d\n", b)
	fmt.Printf("Nilai c: %d\n", c)
	fmt.Printf("Nilai d: %d\n", d)
	fmt.Printf("Nilai e: %d\n", e)


}
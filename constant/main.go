package main

import "fmt"

// membuat CONSTANT di luar fungsi
// CONST tidak bisa deklarasi tanpa assign nilai
const LENGTH int = 10
// const X int
// X = 10

func main() {
	// Membuat constant di dalam fungsi
	const width = 60
	const PI = 3.14
	// PI = 11

	var x = 10
	x = 12

	fmt.Println(LENGTH)
	fmt.Println(width)
	fmt.Println(PI)
	fmt.Println(x)

	/*DEKLARASI MULTIPLE CONSTANT*/
	const (
		X int = 19
		Y			= 3.14
		Z			= "YA"
	)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)

}
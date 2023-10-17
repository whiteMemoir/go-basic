package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	slice1 := mySlice[:3]
	slice2 := mySlice[3:]

	fmt.Println("===[mySlice]===")
	fmt.Println("Data: ", mySlice)
	fmt.Println("Len: ", len(mySlice))
	fmt.Println("Cap: ", cap(mySlice))

	fmt.Println("===[slice1]===")
	fmt.Println("Data: ", slice1)
	fmt.Println("Len: ", len(slice1))
	fmt.Println("Cap: ", cap(slice1))

	fmt.Println("===[slice2]===")
	fmt.Println("Data: ", slice2)
	fmt.Println("Len: ", len(slice2))
	fmt.Println("Cap: ", cap(slice2))

	slice1 = append(slice1, 60)

	fmt.Println("===[Setelah append ke slice1]===")
	fmt.Println("Data slice1: ", slice1)
	fmt.Println("Len slice1: ", len(slice1))
	fmt.Println("Cap slice1: ", cap(slice1))
	fmt.Println("Data slice2: ", slice2)
	fmt.Println("Len slice2: ", len(slice2))
	fmt.Println("Cap slice2: ", cap(slice2))

	slice2 = append(slice2, 70)

	fmt.Println("===[Setelah append ke slice2]===")
	fmt.Println("Data slice1: ", slice1)
	fmt.Println("Len slice1: ", len(slice1))
	fmt.Println("Cap slice1: ", cap(slice1))
	fmt.Println("Data slice2: ", slice2)
	fmt.Println("Len slice2: ", len(slice2))
	fmt.Println("Cap slice2: ", cap(slice2))

}
package main

import (
	"fmt"
	// "log"
)

func InsertionSort(arr ...int) (sorted []int) {
	for i := 1; i < len(arr); i++ {
		currVal := arr[i]
		j := i-1
		for j >= 0 && arr[j] > currVal {
			fmt.Println("ini j ", arr[j], " | ini i", arr[i], arr[j], ">", arr[i],  arr[j] > arr[i])
			arr[j + 1] = arr[j];
			
			j--
		}
		arr[j + 1] = currVal
		fmt.Println("numbers", arr)
	}
	sorted = arr
	return
}

// func InsertionSort(arr ...int) (sorted []int) {
// 	for i := 0; i < len(arr); i++ {
// 		for j := i; j > 0; j-- {
// 			fmt.Println(" dan ini j-1 ", arr[j-1], "ini j ", arr[j], arr[j-1], ">", arr[j],  arr[j-1] > arr[j])
// 			fmt.Println("numbers", arr)
// 			if arr[j-1] > arr[j] {
// 					temp := arr[j]
// 					arr[j] = arr[j-1]
// 					arr[j-1] = temp
// 			}
// 		}
// 	}
// 	sorted = arr
// 	return
// }


func main() {
	// numbers := []int{6,5,3,1,8,7,2,4}
	numbers := InsertionSort(6,5,3,1,8,7,2,4)
	fmt.Println("Setelah disortir", numbers)

}
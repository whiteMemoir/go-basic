package main

import "fmt"

func bubbleSort(arr ...int) (sorted []int) {
	arrLenJ := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arrLenJ; j++ {
			fmt.Println("ini j ", arr[j], " dan ini j+1 ", arr[j+1], arr[j], ">", arr[j+1],  arr[j] > arr[j+1])
			fmt.Println("numbers", arr)
			if arr[j] > arr[j+1]  {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		arrLenJ -= 1
	}
	sorted = arr
	return
}

// func BubbleSort(nums ...int) (sorted []int) {
// 	for i := 0; i < len(nums) - 1; i++ {
// 		for j := i+1; j < len(nums); j++ {
// 			fmt.Println("ini i", i, "dan j ", j,  nums[i], ">", nums[j], nums[i] > nums[j])
// 			// fmt.Println("ini j", j, nums[j])
// 			if nums[i] > nums[j] {
// 				temp := nums[i]
// 				nums[i] = nums[j]
// 				nums[j] = temp
// 			}

// 			fmt.Println(nums)
// 		}
// 	}
// 	return
// }


func main() {
	// numbers := []int{5,6,3,1,8,7,2,4}
	// bubbleSort(numbers...)
	numbers := bubbleSort(5,6,3,1,8,7,2,4)
	fmt.Println("Setelah disortir", numbers)
}
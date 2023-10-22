package main

import "fmt"

func fibonacci(length int) (fibSlice []int){
	if length == 0 {
		return []int{}
	} else if length == 1 {
		return []int{1}
	}

	fibSlice = make([]int, length)
	fibSlice[0], fibSlice[1] = 1, 1

	for i := 2; i < length; i++ {
		fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
	}
	return
}

type Examination struct {
	Name string
	Score int
}

func InsertionSort(arr []Examination) (sorted []Examination, maxScores []int, minScores []int, average float64, median float64) {
	// Menggunakan insertion sort untuk sorting
	for i := 1; i < len(arr); i++ {
		currVal := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Score > currVal.Score {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = currVal
	}

	sum := 0.0
	minScore := arr[0].Score
	maxScore := arr[len(arr)-1].Score

	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i].Score)

		if arr[i].Score == maxScore {
			maxScores = append(maxScores, arr[i].Score)
		}

		if arr[i].Score == minScore {
			minScores = append(minScores, arr[i].Score)
		}
	}

	var medianValue float64
	if len(arr)%2 == 0 {
		medianValue = float64(arr[len(arr)/2-1].Score+arr[len(arr)/2].Score) / 2
	} else {
		medianValue = float64(arr[len(arr)/2].Score)
	}

	average = sum / float64(len(arr))
	sorted = arr
	median = medianValue
	return
}

func Binary(sortedNums []Examination) (above80 []int) {
	target := 80

	low, high := 0, len(sortedNums)-1

	for low <= high {
		middleIndex := (low + high) / 2
		if sortedNums[middleIndex].Score > target {
			for i := middleIndex; i < len(sortedNums); i++ {
				if sortedNums[i].Score > target {
					above80 = append(above80, sortedNums[i].Score)
				} else {
					break
				}
			}
			break
		} else if target < sortedNums[middleIndex].Score {
			high = middleIndex - 1
		} else {
			low = middleIndex + 1
		}
	}
	return above80
}


// func Binary(sortedNums ...Examination) (above80 []int) {
// 	target := 80
// 	if len(sortedNums) <= 1 {
// 		if sortedNums[0].Score == target {
// 			above80 = append(above80, sortedNums[0].Score) 
// 		}
// 		return
// 	}
// 	for {
// 		middleIndex := len(sortedNums) / 2
// 		if sortedNums[middleIndex].Score == target {
// 			above80 = append(above80, sortedNums[0].Score)  
// 		} else if target < sortedNums[middleIndex].Score {
// 			sortedNums = sortedNums[:middleIndex]
// 		} else if target > sortedNums[middleIndex].Score {
// 			sortedNums = sortedNums[middleIndex:]
// 		}
// 		if middleIndex == 0 {
// 			break
// 		}
// 	}
// 	return 
// }

func main() {
	/*===SOAL-1===*/
	fibonacci1 := fibonacci(10)
	fmt.Println(fibonacci1)

	/*===SOAL-2===*/
	examResult := []Examination{
    {
        Name: "Budi",
        Score: 90,
    },
    {
        Name: "Andi",
        Score: 85,
    },
    {
        Name: "Nayla",
        Score: 87,
    },
    {
        Name: "Danu",
        Score: 80,
    },
    {
        Name: "Rahmat",
        Score: 75,
    },
    {
        Name: "Erika",
        Score: 83,
    },
    {
        Name: "Siska",
        Score: 87,
    },
    {
        Name: "Mita",
        Score: 94,
    },
    {
        Name: "Shinta",
        Score: 82,
    },
    {
        Name: "Jojo",
        Score: 83,
    },
	}
	sorted, maxScores, minScores, average, median := InsertionSort(examResult)
	fmt.Println("sorted", sorted)
	fmt.Println("maxScores", maxScores)
	fmt.Println("minScores", minScores)
	fmt.Println("average", average)
	fmt.Println("median", median)

	above80 := Binary(examResult)
	fmt.Println("above80", above80)
}
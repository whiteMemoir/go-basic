package main

import "fmt"

func main() {
	// applicant1 := map[string]interface{}{
	// 	"gender": "female",
	// 	"age": 21,
	// }
	// applicant2 := map[string]interface{}{
	// 	"gender": "male",
	// 	"age": 18,
	// }
		gender1 := "female"
		age1 := 21
		gender2 := "male"
		age2 := 18

		canWork(gender1, age1)
		canWork(gender2, age2)
}

func canWork(gender string, age int) {
	if gender == "female" || age >= 21 {
		fmt.Println("Pelamar dapat dipekerjakan.")
	} else {
		fmt.Println("Pelamar tidak dapat dipekerjakan.")
	}
}
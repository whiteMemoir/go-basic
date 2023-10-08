package main

import "fmt"

func main() {
	vowel1 := "o"
	vowel2 := "A"
	consonant1 := "C"
	consonant2 := "k"

	isVowelOrConsonant(vowel1)
	isVowelOrConsonant(vowel2)
	isVowelOrConsonant(consonant1)
	isVowelOrConsonant(consonant2)
}

func isVowelOrConsonant(char string)  {
	switch char {
	case "a", "A", "i", "I", "u", "U", "e", "E", "o", "O":
		fmt.Printf("Huruf %v adalah huruf vokal.\n", char)
	default:
		fmt.Printf("Huruf %v adalah huruf konsonan.\n", char)
	}
}
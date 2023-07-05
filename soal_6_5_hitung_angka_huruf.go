package main

import (
	"fmt"
	"unicode"
)

func countCharacters(str string) (int, int) {
	var alfabet, digit int

	for _, char := range str {
		if unicode.IsLetter(char) {
			alfabet++
		} else if unicode.IsDigit(char) {
			digit++
		}
	}

	return alfabet, digit
}

func main() {
	varString := "Haloo123"

	alfabet, digit := countCharacters(varString)

	fmt.Printf("Jumlah huruf: %d\n", alfabet)
	fmt.Printf("Jumlah angka: %d\n", digit)
}

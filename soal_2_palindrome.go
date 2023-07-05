package main

import (
	"fmt"
	"strings"
)

func palindrome(text string) bool {
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	text := "kasur ini rusak"
	fmt.Println("Kata/kalimat yang diinput: ", text)
	if palindrome(text) {
		fmt.Println("merupakan palindrom.")
	} else {
		fmt.Println("bukan palindrom.")
	}
}

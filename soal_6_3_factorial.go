package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&num)
	fmt.Printf("Faktorial dari bilangan %d adalah %d", num, factorial(num))
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandom(n int) []int {
	// Membuat slice dengan angka 1 sampai n
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = i + 1
	}

	// Mengacak urutan elemen slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sequence), func(i, j int) {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	})

	return sequence
}

func main() {
	sequence := generateRandom(10)
	for _, num := range sequence {
		fmt.Println(num)
	}
}

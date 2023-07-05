package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 2, 4, 6, 9, 5, 4, 16}

	sort.Ints(nums)

	fmt.Println("bilangan terbesar adalah: ", nums[len(nums)-1])
	fmt.Println("bilangan terkecil adalah: ", nums[0])
}

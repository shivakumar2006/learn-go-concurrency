package main

import "fmt"

func main() {
	nums := []int{3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

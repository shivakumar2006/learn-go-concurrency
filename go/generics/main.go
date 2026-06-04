package main

import "fmt"

func PrintAny[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	PrintAny(items)
}

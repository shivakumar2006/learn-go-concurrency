package main

import "fmt"

func main() {
	counter := counter()
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

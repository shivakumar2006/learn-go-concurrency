package main

import "fmt"

func main() {
	num := 1

	update(&num)

	fmt.Println("outside the function", num)
}

func update(num *int) {
	*num = 5
	fmt.Println("inside the function", *num)
}

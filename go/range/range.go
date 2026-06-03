package main

import "fmt"

func main() {

	// RANGE ON MAPS

	m := map[string]string{"name1": "shiva", "name2": "kumar"}

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	for _, value := range m {
		fmt.Println(value)
	}

	// nums := []int{3, 4, 5}

	// sum := 0

	// for i, num := range nums {
	// 	// sum = sum + num
	// 	fmt.Println(num, i)
	// }

	// fmt.Println(sum)
}

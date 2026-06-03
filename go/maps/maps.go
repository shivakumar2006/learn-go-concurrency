package main

import (
	"fmt"
	"maps"
)

func main() {
	maps1 := map[string]int{"phones": 50, "price": 40}
	maps2 := map[string]int{"phones": 50, "price": 40}

	fmt.Println(maps.Equal(maps1, maps2))

	// val, ok := maps["phones"]

	// fmt.Println(val)

	// if ok {
	// 	fmt.Println("key is present")
	// } else {
	// 	fmt.Println("key is not present")
	// }
}

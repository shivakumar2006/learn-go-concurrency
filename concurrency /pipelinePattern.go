package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	var wg sync.WaitGroup

	// double the value
	wg.Add(1)
	doubleOutput := make(chan int)
	go func() {
		defer wg.Done()
		defer close(doubleOutput)
		for num := range input {
			doubleOutput <- num * 2
		}
	}()

	// square the value
	wg.Add(1)
	squareOutput := make(chan int)
	go func() {
		defer wg.Done()
		defer close(squareOutput)
		for num := range doubleOutput {
			squareOutput <- num * num
		}
	}()

	for result := range squareOutput {
		fmt.Println("Result : ", result)
	}

	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	wg := &sync.WaitGroup{}

// 	data := []int{1, 2, 3, 4, 5}
// 	input := make(chan int, len(data))
// 	for _, d := range data {
// 		input <- d
// 	}
// 	close(input)

// 	wg.Add(1)
// 	doubleOutput := make(chan int)
// 	go func(wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		defer close(doubleOutput)
// 		for num := range input {
// 			doubleOutput <- num * 2
// 		}
// 	}(wg)

// 	wg.Add(1)
// 	squareOutput := make(chan int)
// 	go func(wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		defer close(squareOutput)
// 		for num := range doubleOutput {
// 			squareOutput <- num * num
// 		}
// 	}(wg)

// 	for result := range squareOutput {
// 		fmt.Println("result : ", result)
// 	}

// 	wg.Wait()
// }

// package main

// import "fmt"

// func main() {
// 	// create a initial channel with some data...
// 	data := []int{1, 2, 3, 4, 5}
// 	input := make(chan int, len(data))
// 	for _, d := range data {
// 		input <- d
// 	}
// 	close(input)

// 	// First stage of the pipeline is double the input value
// 	doubleOutput := make(chan int)
// 	go func() {
// 		defer close(doubleOutput)
// 		for num := range input {
// 			doubleOutput <- num * 2
// 		}
// 	}()

// 	// second stage of the pipeline : we square the input value
// 	squareOutput := make(chan int)
// 	go func() {
// 		defer close(squareOutput)
// 		for num := range doubleOutput {
// 			squareOutput <- num * num
// 		}
// 	}()

// 	// third step print the square value
// 	for result := range squareOutput {
// 		fmt.Println("squareValues : ", result)
// 	}
// }

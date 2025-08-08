package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang...")

	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receiver function
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelopen := <-myChannel

		fmt.Println(isChannelopen)
		fmt.Println(val)

		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	// request function
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		close(myChannel)
		// myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {

	// }(myChannel, wg)

	wg.Wait()
}
